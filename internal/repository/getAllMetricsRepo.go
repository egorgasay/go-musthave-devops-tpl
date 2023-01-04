package repository

import "devtool/internal/storage"

func (r Repository) GetAllMetrics() ([]*storage.Metrics, error) {
	return r.repo.GetListOfMetrics()
}
