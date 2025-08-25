package repo

import "github.com/mechiko/dbscan"

func (r *Repository) IsA3() bool {
	if di := r.dbs.Info(dbscan.A3); di != nil {
		return true
	}
	return false
}

func (r *Repository) IsConfig() bool {
	if di := r.dbs.Info(dbscan.Config); di != nil {
		return true
	}
	return false
}

func (r *Repository) IsZnak() bool {
	if di := r.dbs.Info(dbscan.TrueZnak); di != nil {
		return true
	}
	return false
}

func (r *Repository) IsSelf() bool {
	if di := r.dbs.Info(dbscan.Other); di != nil {
		return true
	}
	return false
}
