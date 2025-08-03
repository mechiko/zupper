package views

import (
	"fmt"
	"zupper/domain"
	"zupper/gui/resource"
	"zupper/gui/types"

	// "zupper/gui/views/importutsz"
	// "zupper/gui/views/importvoot"
	// "zupper/gui/views/kontragent"
	// "zupper/gui/views/maintain"
	// "zupper/gui/views/reports"
	// "zupper/gui/views/znak"
	"zupper/gui/views/setup"
	"zupper/repo"
)

var pages = map[string]types.PageConfig{
	"Setup": {Title: "Настройки", Image: "104", NewPage: setup.New, Class: "setup"},
	// "Znak":          {Title: "Коробки", Image: "PngKfh", NewPage: znak.NewPage, Class: "znak"},
	// "ImportTTN":     {Title: "Импорт ТТН", Image: "PngRequester", NewPage: importvoot.NewPage, Class: "importttn"},
	// "ImportTTNUtsz": {Title: "Импорт ТТН", Image: "PngRequester", NewPage: importutsz.NewPage, Class: "importttnutsz"},
	// "Reports":       {Title: "Отчеты", Image: "104", NewPage: reports.NewPage, Class: "reports"},
	// "Utility":       {Title: "Обслуживание", Image: "104", NewPage: maintain.NewPage, Class: "utility"},
	// "Kontragent":    {Title: "Контрагенты", Image: "104", NewPage: kontragent.NewPage, Class: "kontragent"},
}

func CreateTreeMenu(app domain.Apper, repo *repo.Repository) *types.AppmenuTreeModel {
	// var importMenu, trueZnakMenu, reportsMenu *types.AppMenu

	tvm := types.NewAppMenuTreeModel()
	page1 := pages["Setup"]
	setup := tvm.NewRootAppMenu(app, page1, page1.NewPage, imageMenu(app, page1.Image))
	tvm.Menu2NewPage[setup] = page1.NewPage

	// if repo.IsZnak() {
	// 	menu20 := tvm.NewRootMenu("ЧЗ", nil, "103")
	// 	page21 := pages["Znak"]
	// 	trueZnakMenu = menu20.AddChild(app, page21, page21.NewPage, imageMenu(app, page21.Image))
	// 	tvm.Menu2NewPage[trueZnakMenu] = page21.NewPage
	// }
	// if repo.IsA3() {
	// 	menu30 := tvm.NewRootMenu("АлкоХелп3", nil, "103")
	// 	page32 := pages["Reports"]
	// 	reportsMenu = menu30.AddChild(app, page32, page32.NewPage, imageMenu(app, page32.Image))
	// 	tvm.Menu2NewPage[reportsMenu] = page32.NewPage

	// 	page33 := pages["Kontragent"]
	// 	kontragentMenu := menu30.AddChild(app, page33, page33.NewPage, imageMenu(app, page33.Image))
	// 	tvm.Menu2NewPage[kontragentMenu] = page33.NewPage

	// user := app.ImportTTN().User()
	// if user == "voot" {
	// 	page31 := pages["ImportTTN"]
	// 	importMenu = menu30.AddChild(app, page31, page31.NewPage, imageMenu(app, page31.Image))
	// 	tvm.Menu2NewPage[importMenu] = page31.NewPage
	// }
	// if user == "utsz" {
	// 	page31 := pages["ImportTTNUtsz"]
	// 	importMenu = menu30.AddChild(app, page31, page31.NewPage, imageMenu(app, page31.Image))
	// 	tvm.Menu2NewPage[importMenu] = page31.NewPage
	// }
	// }
	// if repo.IsA3() {
	// 	page50 := pages["Utility"]
	// 	utility := tvm.NewRootAppMenu(app, page50, page50.NewPage, imageMenu(app, page50.Image))
	// 	tvm.Menu2NewPage[utility] = page50.NewPage
	// }
	switch app.Options().Application.StartPage {
	case "setup":
		tvm.SetDefaultMenu(setup)
	default:
		tvm.SetDefaultMenu(setup)
	}
	return tvm
}

func imageMenu(app domain.Apper, img string) interface{} {
	if img == "" {
		return ""
	}
	if icon, err := resource.New(app).Icon(img); err != nil {
		app.Logger().Errorf("guiwalk:treemenu image[%s] %s", img, err.Error())
		panic(fmt.Sprintf("guiwalk:treemenu image[%s] %s", img, err.Error()))
	} else {
		return icon
	}
}
