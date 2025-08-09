package repo

func (r *Repository) ConfigDBPing() (result bool) {
	defer func() {
		if rec := recover(); rec != nil {
			result = false
		}
	}()
	if err := r.ConfigDB().Ping(); err != nil {
		r.dbs.ConfigInfo().Exists = false
		return false
	}
	r.dbs.ConfigInfo().Exists = true
	return true
}

func (r *Repository) A3DBPing() (result bool) {
	defer func() {
		if rec := recover(); rec != nil {
			result = false
		}
	}()
	if err := r.A3DB().Ping(); err != nil {
		r.dbs.A3().Exists = false
		return false
	}
	r.dbs.A3().Exists = true
	return true
}

func (r *Repository) ZnakDBPing() (result bool) {
	defer func() {
		if rec := recover(); rec != nil {
			result = false
		}
	}()
	if err := r.ZnakDB().Ping(); err != nil {
		r.dbs.Znak().Exists = false
		return false
	}
	r.dbs.Znak().Exists = true
	return true
}

func (r *Repository) SelfDBPing() (result bool) {
	defer func() {
		if rec := recover(); rec != nil {
			result = false
		}
	}()
	if err := r.Self().Ping(); err != nil {
		r.dbs.Self().Exists = false
		return false
	}
	r.dbs.Self().Exists = true
	return true
}
