package repo

import (
	"errors"
	"fmt"

	"github.com/mechiko/dbscan"
)

// при инициализации приложения этот метод вызывается однажды и прописывает объект доступа
// к базе данных, далее проверяет версию БД возможна ошибка и нужно выходить из приложения
func (r *Repository) prepareSelf() (err error) {
	defer func() {
		if rr := recover(); rr != nil {
			err = fmt.Errorf("%s panic %v", modError, rr)
		}
	}()
	selfInfo := r.Info(dbscan.Other)
	self, err := selfInfo.Connect()
	if err != nil {
		return fmt.Errorf("%s self connext error %v", modError, err)
	}
	if self != nil {
		defer func() {
			if errClose := self.Close(); errClose != nil {
				err = errors.Join(err, errClose)
			}
		}()
	}

	info := r.Info(dbscan.Other)
	if info == nil {
		return fmt.Errorf("%s get info: nil", modError)
	}
	if !info.Exists {
		return fmt.Errorf("%s get self db error!", modError)
	}
	return nil
}
