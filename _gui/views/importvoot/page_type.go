package importvoot

import (
	"zupper/gui/types"

	"github.com/mechiko/walk"
)

const modError = "gui:view:importttn"

// A4 width:794px; height:1122px;
type ImportTTNPage struct {
	*walk.Composite
	app  types.IApp
	form walk.Form
	db   *walk.DataBinder

	lblProccessingError *walk.Label
	stageError          *walk.Composite
	lblError            *walk.Label
	lblMsg              *walk.Label
	lblFile             *walk.Label
	lblCountTtn         *walk.Label
	btnSrc              *walk.PushButton
	btnCancel           *walk.PushButton
	btnExamen           *walk.PushButton
	btnImportTtn        *walk.PushButton
	btnControlTtn       *walk.PushButton
	// fifoCheckBox       *walk.CheckBox
	// splitCheckBox      *walk.CheckBox
	reimportCheckBox *walk.CheckBox
	// ignoreRestCheckBox *walk.CheckBox
	// filterBottlingFld  *walk.LineEdit
	// startBottlingFld   *walk.LineEdit
	emptyTtnCheckBox *walk.CheckBox
}
