package adminka

import (
	"zupper/domain"
)

const modError = "pkg:usecase"

type adminka struct {
	domain.Apper
}

// New -.
func New(a domain.Apper) *adminka {
	return &adminka{
		Apper: a,
	}
}
