package uctemplate

import (
	"testing"
	"zupper/domain"

	"github.com/stretchr/testify/assert"
)

var testsParse = []struct {
	name     string
	err      bool
	template string
	model    interface{}
}{
	// the table itself
	{"test 0", true, "tmplAdminReport.html", &domain.AdminReport{}},
}

func TestTemplate(t *testing.T) {
	// The execution loop
	// Capture tt for safety, use NoError, and put expected before actual in Equal
	for _, tt := range testsParse {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			templ := NewTemplate("", false)
			// вызов шаблона в него передаем имя шаблона как имя файла шаблона
			_, err := templ.tmplHtml(tmplAdminReport, tt.template, tt.model, nil)
			if tt.err {
				assert.NotNil(t, err, "ожидаем ошибку")
			} else {
				// ожидаем отсутствие ошибки
				assert.NoError(t, err)
			}
		})
	}
}
