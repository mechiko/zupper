package znaktools

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
func (p *ZnakToolsPage) SetSendFunc(f func(domain.Model)) {
	p.sendChan = f
}

// обновляет по модели страницу
func (p *ZnakToolsPage) Update() {
}

func (p *ZnakToolsPage) Clear() {
}
