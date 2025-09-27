package a3

import (
	"errors"
	"fmt"
	"zupper/domain"

	"github.com/upper/db/v4"
)

func (a *DbA3) PartnerByFsrarId(fsrarid string) (ap *domain.PartnerEgais, err error) {
	ap = &domain.PartnerEgais{}
	coll := a.dbSession.Collection("partners_egais")
	if err := coll.Find(db.Cond{"client_reg_id": fsrarid}).One(ap); err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			return nil, fmt.Errorf("%s: partner %q not found", modError, fsrarid)
		}
		return nil, fmt.Errorf("%s: partners_egais %q: %w", modError, fsrarid, err)
	}
	return ap, nil
}
