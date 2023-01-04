package mapstorage

import (
	"devtool/internal/storage"
	"errors"
)

func (ms *MapStorage) UpdateOneMetric(mt *storage.Metrics) (count float64, err error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if mt.MType == "gauge" {
		ms.store[mt.ID] = *mt.Value
	} else if mt.MType == "counter" {
		ms.store[mt.ID] = float64(*mt.Delta) + ms.store[mt.ID]
	} else {
		return 0, errors.New("unsupported metric type")
	}

	return ms.store[mt.ID], nil
}
