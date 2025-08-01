package kontragent

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"zupper/domain"
	"zupper/usecase/services"
	"zupper/utility"
)

func (p *KontragentListPage) ExportExcelOborot(kontragent *domain.PartnersOrigin) error {
	if file, err := services.New(p.app).PartnerOborotExcelFile(kontragent, ""); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	} else {
		p.app.Logger().Debugf("ExportExcelOborot file %s", file)
	}
	return nil
}

func (p *KontragentListPage) ExporеPartnerOborot(kontragent *domain.PartnersOrigin) error {
	ext := "xml"
	fileName := strings.ReplaceAll(utility.ClearForFileName(kontragent.ClientShortName), " ", "_")
	startDate := fmt.Sprintf("%02d.%02d.%4d", p.app.StartDate().Day(), p.app.StartDate().Month(), p.app.StartDate().Year())
	endDate := fmt.Sprintf("%02d.%02d.%4d", p.app.EndDate().Day(), p.app.EndDate().Month(), p.app.EndDate().Year())
	fileName = fmt.Sprintf("%s %s %s %s Обороты %s %s.%s", fileName, kontragent.ClientRegID, kontragent.ClientInn, kontragent.ClientKPP, startDate, endDate, ext)
	if kontragent.ClientKPP == "" {
		fileName = fmt.Sprintf("%s %s %s Обороты %s %s.%s", fileName, kontragent.ClientRegID, kontragent.ClientInn, startDate, endDate, ext)
	}
	fileName = filepath.Join(p.app.Output(), fileName)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func() error {
		if err := file.Close(); err != nil {
			return fmt.Errorf("%w", err)
		}
		return nil
	}()

	if xmlStr, err := services.New(p.app).XmlPartnerOborot(kontragent.ClientRegID); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	} else {
		if _, err := file.WriteString(xmlStr); err != nil {
			return fmt.Errorf("%s %w", modError, err)
		}
	}
	return nil
}

func (p *KontragentListPage) ExportFirmShipper(kontragent *domain.PartnersOrigin) error {
	ext := "xml"
	fileName := strings.ReplaceAll(utility.ClearForFileName(kontragent.ClientShortName), " ", "_")
	startDate := fmt.Sprintf("%02d.%02d.%4d", p.app.StartDate().Day(), p.app.StartDate().Month(), p.app.StartDate().Year())
	endDate := fmt.Sprintf("%02d.%02d.%4d", p.app.EndDate().Day(), p.app.EndDate().Month(), p.app.EndDate().Year())
	fileName = fmt.Sprintf("%s %s %s %s Организации %s %s.%s", fileName, kontragent.ClientRegID, kontragent.ClientInn, kontragent.ClientKPP, startDate, endDate, ext)
	if kontragent.ClientKPP == "" {
		fileName = fmt.Sprintf("%s %s %s Организации %s %s.%s", fileName, kontragent.ClientRegID, kontragent.ClientInn, startDate, endDate, ext)
	}
	fileName = filepath.Join(p.app.Output(), fileName)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func() error {
		if err := file.Close(); err != nil {
			return fmt.Errorf("%w", err)
		}
		return nil
	}()

	if xmlStr, err := services.New(p.app).XmlProducerShipperByPartnerFile(kontragent.ClientRegID); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	} else {
		if _, err := file.WriteString(xmlStr); err != nil {
			return fmt.Errorf("%s %w", modError, err)
		}
	}
	return nil
}
