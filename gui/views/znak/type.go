package znak

import (
	"fmt"
	"zupper/domain"
	"zupper/domain/znakagregate"
	"zupper/gui/types"
	"zupper/reductor"
	"zupper/repo"

	"github.com/mechiko/walk"
)

const modError = "gui:view:znak"

type ZnakPage struct {
	*walk.Composite
	domain.Apper
	model    domain.Model
	sendChan func(domain.Model)

	parent    walk.Form
	smallFont *walk.Font
	tableFont *walk.Font

	groupLbl       *walk.Label
	packageLbl     *walk.Label
	ipsCombo       *walk.ComboBox
	groupItogLbl   *walk.Label
	packageItogLbl *walk.Label
	fileLbl        *walk.Label
	filePb         *walk.PushButton
	filePbA3       *walk.PushButton
	filePbXml      *walk.PushButton
	filePb1C       *walk.PushButton
	filePbCsv      *walk.PushButton
	waitStateLbl   *walk.Label
}

func New(parent walk.Container, app domain.Apper, repo *repo.Repository) (pp types.Page, err error) {
	// func NewPage(parent walk.Container, app domain.Apper) (pp types.Page, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s NewPage ZnakPage panic %v", modError, r)
		}
	}()
	p := &ZnakPage{
		Apper:  app,
		parent: parent.Form(),
		model:  domain.ZnakAgregate,
	}
	p.smallFont, _ = walk.NewFont("JetBrains Mono", 9, 0)
	p.tableFont, _ = walk.NewFont("JetBrains Mono", 10, walk.FontBold)

	var model *znakagregate.ZnakAgregate
	// инициализируем модель и сохраняем в редукторе
	// если таковой еще нет в нем, предохраняется модель уже созданная и рабочая
	if reductor.Instance().IsExistModel(p.model) {
		model, err = p.Model()
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	} else {
		if model, err = znakagregate.New(p.Apper, repo); err != nil {
			return nil, fmt.Errorf("view:znak ошибка создания модели %s %v", p.model, err)
		}
		if err := reductor.Instance().SetModel(model, false); err != nil {
			return nil, fmt.Errorf("view:znak ошибка записи модели в редуктор %s %v", p.model, err)
		}
	}
	if err = p.dclCreate(parent, model); err != nil {
		return nil, fmt.Errorf("page znak dcl create %w", err)
	}
	return p, nil
}
