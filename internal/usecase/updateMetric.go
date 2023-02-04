package usecase

import (
	"devtool/internal/storage"
	"encoding/json"
	"fmt"
)

func (uc UseCase) UpdateMetricByJSON(b []byte) ([]byte, error) {
	var metrics storage.Metrics
	err := json.Unmarshal(b, &metrics)
	if err != nil {
		return nil, fmt.Errorf("UpdateMetricByJSON: %w", err)
	}

	count, err := uc.service.DB.UpdateMetric(&metrics)
	if err != nil {
		return nil, fmt.Errorf("UpdateMetricByJSON: %w", ErrNotFound)
	}

	if metrics.MType == "gauge" {
		metrics.Value = count
	} else {
		metrics.Delta = int64(count)
	}

	metrics.UpdateNeeded = true

	byteJSON, err := json.MarshalIndent(metrics, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("UpdateMetricByJSON: %w", ErrNotFound)
	}

	return byteJSON, nil
}

func (uc UseCase) UpdateMetric(val float64, metricType string, name string) error {
	var mt *storage.Metrics

	switch metricType {
	case "gauge":
		mt = &storage.Metrics{
			ID:    name,
			MType: metricType,
			Value: val,
		}
	case "counter":
		mt = &storage.Metrics{
			ID:    name,
			MType: metricType,
			Delta: int64(val),
		}
	}

	_, err := uc.service.DB.UpdateMetric(mt)
	if err != nil {
		return err
	}

	return nil
}
