package mainwindow

// нажатие кнопки на вкл сканирования
func (w *MainWindow) SbiScanPress() {
	// тут надо отправить в редуктор команду смена Скан
	// w.SetScanUTM(!w.Reductor().Model().Gui.MainWindow.StatusBar.Scan)
}

func (w *MainWindow) clickHistoryState() {
	// тут надо отправить в редуктор команду
}

func (w *MainWindow) SbiLicensePress() {
	// тут надо отправить в редуктор команду
	// msg := entity.Message{
	// 	Sender: "mainwindow.SbiLicensePress",
	// 	Cmd:    "license",
	// 	Model:  nil,
	// }
	// w.Effects().ChanIn() <- msg
	// w.Logger().Debugf("%s effects chanin msg %s", modError, msg.Cmd)
}

func (w *MainWindow) SbiUtmPress() {
	// тут надо отправить в редуктор команду
	// msg := entity.Message{
	// 	Sender: "mainwindow.SbiUtmPress",
	// 	Cmd:    "utm",
	// 	Model:  nil,
	// }
	// w.Effects().ChanIn() <- msg
	// w.Logger().Debugf("%s effects chanin msg %s", modError, msg.Cmd)
}

func (w *MainWindow) SbiFsrarIdPress() {
	// тут надо отправить в редуктор команду
}
