package repository

import (
	_ "github.com/mattn/go-sqlite3"
)

func (r Repository) GetMetric(name string) (float64, error) {
	return r.repo.GetOneMetric(name)
}
