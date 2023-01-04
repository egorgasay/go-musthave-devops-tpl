package mapstorage

import "errors"

func (ms *MapStorage) GetOneMetric(name string) (val float64, err error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	var ok bool
	if val, ok = ms.store[name]; !ok {
		return 0, errors.New("can't get metric")
	}

	return val, nil
}
