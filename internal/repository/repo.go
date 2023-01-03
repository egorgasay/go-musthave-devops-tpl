package repository

import (
	"database/sql"
)

type Config struct {
	DriverName     string
	DataSourceName string
}

type IMemStorage interface {
	Ping() error
	Close() error
	Query(string, ...any) (*sql.Rows, error)
	Exec(string, ...any) (sql.Result, error)
	QueryRow(string, ...any) *sql.Row
}

type MemStorage struct {
	DB IMemStorage
}

func NewMemStorage(cfg *Config) (*MemStorage, error) {
	db, err := sql.Open(cfg.DriverName, cfg.DataSourceName)
	if err != nil {
		return nil, err
	}

	return &MemStorage{
		DB: db,
	}, nil
}

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

type MemStorageMethods struct {
	DB IMemStorage
}

//func NewMemStorageMethods(ms IMemStorage) *MemStorageMethods {
//	return &MemStorageMethods{DB: ms}
//}
