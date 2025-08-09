package gui

import (
	"fmt"
	"zupper/domain/models/statusbar"
	"zupper/reductor"
)

func (s *guiService) StatusBarInit() error {
	statusModel, err := statusbar.New(s, s.repo)
	if err != nil {
		return fmt.Errorf("ошибка создания модели statusbar %w", err)
	}
	if err := reductor.Instance().SetModel(statusModel, false); err != nil {
		return fmt.Errorf("ошибка редуктора сохранения модели statusbar %w", err)
	}
	return nil
}
