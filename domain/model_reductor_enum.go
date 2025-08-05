package domain

type Model string

const (
	Application Model = "application"
	TrueClient  Model = "trueclient"
	StatusBar   Model = "statusbar"
)

func IsValidModel(s string) bool {
	switch Model(s) {
	case Application, TrueClient, StatusBar:
		return true
	default:
		return false
	}
}
