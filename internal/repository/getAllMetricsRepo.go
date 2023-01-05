package repository

import "devtool/internal/storage"

func (r *Repository) GetAllMetrics() ([]*storage.Metrics, error) {
	r.Mu.RLock()
	defer r.Mu.RUnlock()

	metrics := make([]*storage.Metrics, 0, len(r.Store))

	for key, value := range r.Store {
		val := value
		metric := &storage.Metrics{
			ID:    key,
			Value: &val,
		}

		metrics = append(metrics, metric)
	}

	return metrics, nil
}
