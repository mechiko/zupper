package znakpacks

import (
	"fmt"
	"zupper/domain"
	"zupper/domain/models/znakagregate"
	"zupper/gui/types"
	"zupper/reductor"
	"zupper/repo"

	"github.com/mechiko/walk"
)

const modError = "gui:view:znakpacks"

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
	fileLblCn      *walk.Label
	fileLblXml     *walk.Label
	fileLblA3      *walk.Label
	fileLbl1C      *walk.Label
	fileLblCsv     *walk.Label

	filePb       *walk.PushButton
	filePbA3     *walk.PushButton
	filePbXml    *walk.PushButton
	filePb1C     *walk.PushButton
	filePbCsv    *walk.PushButton
	waitStateLbl *walk.Label
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
	p.smallFont = p.initFont("JetBrains Mono", 9, 0)
	p.tableFont = p.initFont("JetBrains Mono", 10, walk.FontBold)

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
			return nil, fmt.Errorf("view:znakpacks ошибка создания модели %s %v", p.model, err)
		}
		if err := reductor.Instance().SetModel(model, false); err != nil {
			return nil, fmt.Errorf("view:znakpacks ошибка записи модели в редуктор %s %v", p.model, err)
		}
	}
	if err = p.dclCreate(parent, model); err != nil {
		return nil, fmt.Errorf("page znakpacks dcl create %w", err)
	}
	return p, nil
}

func (p *ZnakPage) initFont(name string, size int, style walk.FontStyle) *walk.Font {
	f, err := walk.NewFont(name, size, style)
	if err != nil && p.Logger() != nil {
		p.Logger().Warnf("%s font create (size %d): %v", modError, size, err)
	}
	return f
}
