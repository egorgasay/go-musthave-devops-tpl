package repository

import (
	"devtool/internal/storage"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

func (r *Repository) UpdateMetric(mt *storage.Metrics) (count float64, err error) {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	if mt.MType == "gauge" {
		r.Store[mt.ID] = mt.Value
	} else if mt.MType == "counter" {
		r.Store[mt.ID] = float64(mt.Delta) + r.Store[mt.ID]
	} else {
		return 0, errors.New("unsupported metric type")
	}

	storage.StorageRelevance.Mu.Lock()
	storage.StorageRelevance.UpdateNeeded[mt.ID] = struct{}{}
	storage.StorageRelevance.Status = false
	storage.StorageRelevance.Mu.Unlock()

	return r.Store[mt.ID], nil
}
