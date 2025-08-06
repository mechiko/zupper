package reports

import (
	"zupper/domain"
)

// обязательные методы для реализации интерфейса types.Page

// через эту функцию прописывает метод обратного вызова для записи в канал смены состояния
// так можно использовать в кнопке вызов отправки в канал
//
//	if p.sendChan != nil {
//		p.sendChan(p.model)
//	}
func (p *ReportsPage) SetSendFunc(f func(domain.Model)) {
	p.sendChan = f
}

// обновляет по модели страницу
func (p *ReportsPage) Update() {
	model, err := p.Model()
	if err != nil {
		p.Logger().Errorf("view:setup update error %v", err)
	}
	p.start.SetDate(model.StartDate())
	p.end.SetDate(model.EndDate())
}

func (p *ReportsPage) Clear() {
}
