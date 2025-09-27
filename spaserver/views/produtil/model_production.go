package produtil

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"zupper/domain"
	"zupper/domain/models/application"
	"zupper/reductor"
	"zupper/repo"

	"github.com/mechiko/utility"
	"github.com/upper/db/v4"
)

type PrdReport struct {
	domain.Apper
	repo      *repo.Repository
	model     *application.Application
	Report    *domain.ProductionReport
	Products  []*domain.ProductionProduct
	MapCodeAP map[string]*domain.ApEgais
}

func NewPrdReport(app domain.Apper) (prdr *PrdReport, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	rp, err := repo.GetRepository()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	dbA3, err := rp.LockA3()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer func() {
		if cerr := rp.UnlockA3(dbA3); cerr != nil {
			err = errors.Join(err, cerr)
		}
	}()

	model, err := reductor.Instance().Model(domain.Application)
	if err != nil {
		return nil, fmt.Errorf("get reductor model domain.Application %w", err)
	}
	mdl, ok := model.(*application.Application)
	if !ok {
		return nil, fmt.Errorf("model wrong type %T %w", model, err)
	}
	// справочник кодов АП
	aps, err := dbA3.CodeApMap()
	if err != nil {
		return nil, fmt.Errorf("get map code ap error %w", err)
	}
	prdr = &PrdReport{
		Apper:     app,
		repo:      rp,
		model:     mdl,
		MapCodeAP: aps,
		Products:  make([]*domain.ProductionProduct, 0),
	}

	prdr.productionReport()
	owner, err := dbA3.PartnerByFsrarId(mdl.FsrarID)
	if err != nil {
		return nil, fmt.Errorf("model get owner error %w", err)
	}
	prdr.setProducer(owner)
	return prdr, err
}

func (prd *PrdReport) scanDayUtilisation(day map[string]int) error {
	for codeAp, quantity := range day {
		codeEgais, exist := prd.MapCodeAP[codeAp]
		if !exist {
			return fmt.Errorf("%s не найден", codeAp)
		}
		productQuantity := float64(0)
		volume, err := strconv.ParseFloat(codeEgais.ProductCapacity, 64)
		if err != nil {
			return fmt.Errorf("%s ошибка конвертации объема тары", codeEgais.ProductCapacity)
		}
		switch strings.ToLower(codeEgais.ProductUnitType) {
		case "packed":
			productQuantity = float64(quantity)
		case "unpacked":
			// далы = емкость тары * количество марок отчета нанесения / 10
			productQuantity = float64(quantity) * volume / 10
		default:
			return fmt.Errorf("%s ошибочный тип упаковки", codeEgais.ProductUnitType)
		}
		prd.AddProduct(codeEgais, fmt.Sprintf("%0.4f", productQuantity))
	}
	return nil
}

func (prd *PrdReport) productionReport() {
	identity := fmt.Sprintf("ah3-rppg-%s", utility.String(16))
	prRepVersion := prd.Options().Application.ProduceReportVersion
	if prRepVersion == "" {
		prRepVersion = "6"
	}
	prd.Report = &domain.ProductionReport{
		CreateDate:          time.Now().Format(prd.Options().Layouts.TimeLayoutClear),
		DocIdentity:         identity,
		DocType:             "Производство",
		DocNumber:           "",
		DocDate:             time.Now().Format(prd.Options().Layouts.TimeLayoutDay),
		DocProducedDate:     "",
		DocComment:          "",
		ProducerType:        "",
		ProducerClientRegId: "",
		ProducerInn:         "",
		ProducerKpp:         "",
		ProducerFullName:    "",
		ProducerShortName:   "",
		ProducerCountryCode: "",
		ProducerRegionCode:  "",
		ProducerDescription: "",
		Version:             prRepVersion,
		State:               "Создан",
		Status:              "Не проведён",
		ReplyId:             "",
		Archive:             0,
		Xml:                 "",
	}
}

func (prd *PrdReport) setProducer(pr *domain.PartnerEgais) {
	prd.Report.ProducerType = pr.ClientType
	prd.Report.ProducerClientRegId = pr.ClientRegId
	prd.Report.ProducerInn = pr.ClientInn
	prd.Report.ProducerKpp = pr.ClientKpp
	prd.Report.ProducerFullName = pr.ClientFullName
	prd.Report.ProducerShortName = pr.ClientShortName
	prd.Report.ProducerCountryCode = pr.ClientCountryCode
	prd.Report.ProducerRegionCode = pr.ClientRegionCode
	prd.Report.ProducerDescription = pr.ClientDescription
}

// добавляем строку в отчет производства
func (prd *PrdReport) AddProduct(pr *domain.ApEgais, quantity string) {
	idnt := fmt.Sprintf("%d", len(prd.Products)+1)
	r := &domain.ProductionProduct{
		IdProductionReports: prd.Report.Id,
		ProductFullName:     pr.ProductFullName,
		ProductCapacity:     pr.ProductCapacity,
		ProductAlcVolume:    pr.ProductAlcVolume,
		ProductAlcVolumeMin: "0.000",
		ProductAlcVolumeMax: "0.000",
		ProductAlcCode:      pr.ProductAlcCode,
		ProductCode:         pr.ProductCode,
		ProductUnitType:     pr.ProductUnitType,
		ProductIdentity:     idnt,
		ProductQuantity:     quantity,
		ProductParty:        "1",
		ProductComment:      "",
		ProducerType:        pr.ProducerType,
		ProducerClientRegId: pr.ProducerClientRegId,
		ProducerInn:         pr.ProducerInn,
		ProducerKpp:         pr.ProducerKpp,
		ProducerFullName:    pr.ProducerFullName,
		ProducerShortName:   pr.ProducerShortName,
		ProducerCountryCode: pr.ProducerCountryCode,
		ProducerRegionCode:  pr.ProducerRegionCode,
		ProducerDescription: pr.ProducerDescription,
	}
	prd.Products = append(prd.Products, r)
}

// запись отчета нанесения
func (prd *PrdReport) WriteBD() (err error) {
	dbA3, err := prd.repo.LockA3()
	if err != nil {
		return fmt.Errorf("LockA3: %w", err)
	}
	defer func() {
		if uerr := prd.repo.UnlockA3(dbA3); uerr != nil {
			err = errors.Join(err, uerr)
		}
	}()
	return dbA3.Sess().Tx(func(tx db.Session) error {
		return prd.writeReport(tx)
	})
}

func (prd *PrdReport) writeReport(tx db.Session) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	if err := tx.Collection("production_reports").InsertReturning(prd.Report); err != nil {
		return err
	} else {
		for _, product := range prd.Products {
			product.IdProductionReports = prd.Report.Id
			if _, err := tx.Collection("production_products").Insert(product); err != nil {
				return err
			}
		}
	}
	return nil
}
