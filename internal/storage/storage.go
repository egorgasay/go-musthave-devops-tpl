package storage

import (
	"sync"
)

//type IStorage interface {
//	GetListOfMetrics() ([]*Metrics, error)
//	GetOneMetric(name string) (float64, error)
//	UpdateOneMetric(mt *Metrics) (count float64, err error)
//	Restore() error
//	StageChanges() error
//}
//
//type Storage struct {
//	IStorage
//}

type Row struct {
	MName  string
	MValue float64
}

type Relevance struct {
	UpdateNeeded map[string]struct{}
	Status       bool
	Mu           sync.RWMutex
}

var StorageRelevance = Relevance{
	UpdateNeeded: make(map[string]struct{}),
	Status:       true,
}

type IBackupStorage interface {
	GetAllMetrics() ([]*Metrics, error)
	DoBackup(ms map[string]float64) error
}

type BackupStorage struct {
	IBackupStorage
}

type Metrics struct {
	ID           string  `json:"id"`              // имя метрики
	MType        string  `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta        int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value        float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
	UpdateNeeded bool    `json:"-"`
	Hash         string  `json:"hash,omitempty"` // значение хеш-функции
}
