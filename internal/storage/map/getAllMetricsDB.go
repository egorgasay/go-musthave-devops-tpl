package mapstorage

import "devtool/internal/storage"

func (ms *MapStorage) GetListOfMetrics() ([]*storage.Metrics, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	metrics := make([]*storage.Metrics, 0, len(ms.store))

	for key, value := range ms.store {
		val := value
		metric := &storage.Metrics{
			ID:    key,
			Value: &val,
		}

		metrics = append(metrics, metric)
	}

	return metrics, nil
}
