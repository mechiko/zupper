package mainwindow

// func (w *MainWindow) ReductorUpdaterGUI(cmd string, model entity.Model) (cmdOut string, modelNew entity.Model) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			w.Logger().Errorf("%s ReductorUpdaterGUI %s panic %v", modError, cmd, r)
// 		}
// 	}()
// 	// обновляем по модели текущую страницу
// 	switch cmd {
// 	case "reload":
// 		w.Synchronize(func() {
// 			t := views.CreateTreeMenu(w)
// 			if err := w.tv.SetModel(t); err != nil {
// 				w.Logger().Error(err.Error())
// 			}
// 			w.Tvm = t
// 			defaultMenu := w.Tvm.DefaultMenu()
// 			if err := w.tv.SetCurrentItem(defaultMenu); err != nil {
// 				w.Logger().Error(err.Error())
// 			}
// 			w.сhangePage()
// 		})
// 	case "title":
// 		w.Synchronize(func() {
// 			w.SetTitle(model.Gui.MainWindow.Title)
// 		})
// 	case "statusbar":
// 		w.Synchronize(func() {
// 			w.UpdateStatusBar(model)
// 		})
// 	case "all":
// 		w.Synchronize(func() {
// 			w.SetTitle(model.Gui.MainWindow.Title)
// 		})
// 		w.Synchronize(func() {
// 			w.UpdateStatusBar(model)
// 		})
// 		page := w.Tvm.CurrentPage()
// 		if page != nil {
// 			page.Synchronize(func() {
// 				page.ReductorChangeState(model)
// 			})
// 		}
// 	case "page":
// 		page := w.Tvm.CurrentPage()
// 		if page != nil {
// 			page.Synchronize(func() {
// 				page.ReductorChangeState(model)
// 			})
// 		}
// 	}
// 	return "", model
// }
