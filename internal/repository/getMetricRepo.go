package repository

import (
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

func (r *Repository) GetMetric(name string) (val float64, err error) {
	r.Mu.RLock()
	defer r.Mu.RUnlock()

	var ok bool
	if val, ok = r.Store[name]; !ok {
		return 0, errors.New("can't get metric")
	}

	return val, nil
}
