package usecase

import (
	"devtool/internal/storage"
	"encoding/json"
	"errors"
	"strconv"
)

func (uc UseCase) GetMetric(name string) (string, error) {
	val, err := uc.service.DB.GetMetric(name)
	if err != nil {
		return "", err
	}

	return strconv.FormatFloat(val, 'f', -1, 64), nil
}

func (uc UseCase) GetMetricByJSON(b []byte) ([]byte, error) {
	var metric storage.Metrics

	err := json.Unmarshal(b, &metric)
	if err != nil {
		return nil, err
	}

	val, err := uc.service.DB.GetMetric(metric.ID)
	if err != nil {
		return nil, ErrNotFound
	}

	if metric.MType == "gauge" {
		metric.Value = val
	} else if metric.MType == "counter" {
		metric.Delta = int64(val)
	} else {
		return nil, errors.New("not implemented")
	}

	outputJSON, err := json.Marshal(metric)
	if err != nil {
		return nil, nil
	}

	return outputJSON, nil
}

func (uc UseCase) GetAllMetrics() ([]*storage.Metrics, error) {
	metrics, err := uc.service.DB.GetAllMetrics()
	if err != nil {
		return nil, err
	}

	return metrics, nil
}
