package trueclient

import (
	"zupper/domain"
	"zupper/reductor"
)

// сохраняем в конфиге и в модели токены и время получения
func (t *trueClient) Save(model *TrueClientModel) {
	model.AuthTime = t.authTime
	model.TokenGIS = t.tokenGis
	model.TokenSUZ = t.tokenSuz
	reductor.Instance().SetModel(domain.TrueClient, model)
}
