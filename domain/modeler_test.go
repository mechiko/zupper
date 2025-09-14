package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelFromText(t *testing.T) {
	model, err := ModelFromString("123")
	assert.NotNil(t, err, "ожидаем ошибку")
	fmt.Println(model)
}
