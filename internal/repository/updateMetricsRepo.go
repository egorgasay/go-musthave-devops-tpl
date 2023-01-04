package repository

import (
	"devtool/internal/storage"

	_ "github.com/mattn/go-sqlite3"
)

func (r Repository) UpdateMetric(mt *storage.Metrics) (count float64, err error) {
	return r.repo.UpdateOneMetric(mt)
}
