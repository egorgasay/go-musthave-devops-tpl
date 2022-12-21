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
	Type  string
	Name  string
	Value float64
}
type MemStorageMethods struct {
	DB IMemStorage
}

//func NewMemStorageMethods(ms IMemStorage) *MemStorageMethods {
//	return &MemStorageMethods{DB: ms}
//}
