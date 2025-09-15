package znaktools

import (
	"zupper/domain"
)

// Required methods to implement the types.Page interface

// SetSendFunc sets the callback method for writing to the state change channel
// This allows button calls to send to the channel:
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

// Clear resets the page to its initial state
func (p *ZnakToolsPage) Clear() {
	// TODO: Implement page clearing logic
}
