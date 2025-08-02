package importvoot

import (
	"path/filepath"

	"zupper/entity"
)

// func NewPage(parent walk.Container, app types.IApp) (pp types.Page, err error)
func (p *ImportTTNPage) init() error {
	PageData = new(pageData)
	PageData.Browser = p.app.Configuration().Browser
	PageData.Pwd = p.app.Pwd()
	PageData.Export = p.app.Export()
	PageData.Debug = entity.Mode == "development"
	if PageData.Browser == "" {
		PageData.Browser = "по умолчанию"
	}
	PageData.File = "[выберите файл для обработки]"
	PageData.Input = p.app.Configuration().Import.Input
	PageData.Output = filepath.Join(p.app.Pwd(), p.app.Configuration().Output)
	PageData.Err = ""
	PageData.Message = ""

	return nil
}
