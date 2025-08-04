package domain

type Model string

const (
	Setup       Model = "setup"
	Application Model = "application"
	TrueClient  Model = "trueclient"
	StatusBar   Model = "statusbar"
)

func IsValidModel(s string) bool {
	switch Model(s) {
	case Setup, Application, TrueClient, StatusBar:
		return true
	default:
		return false
	}
}
