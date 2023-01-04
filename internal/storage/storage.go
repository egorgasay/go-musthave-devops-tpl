package storage

type IStorage interface {
	GetListOfMetrics() ([]*Metrics, error)
	GetOneMetric(name string) (float64, error)
	UpdateOneMetric(mt *Metrics) (count float64, err error)
}

type Storage struct {
	IStorage
}

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}
