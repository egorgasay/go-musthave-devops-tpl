package repository

import (
	"database/sql"
	"devtool/internal/storage"
	dbstorage "devtool/internal/storage/db"
	filestorage "devtool/internal/storage/file"
	mapstorage "devtool/internal/storage/map"
	"os"
	"time"
)

type Config struct {
	DriverName     string
	DataSourceName string
	TimeBeforeSave time.Duration
}

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

type Repository struct {
	repo storage.IStorage
}

func New(cfg *Config) (*Repository, error) {
	if cfg == nil {
		panic("конфигурация задана некорректно")
	}

	switch cfg.DriverName {
	case "sqlite3":
		db, err := sql.Open(cfg.DriverName, cfg.DataSourceName)
		if err != nil {
			return nil, err
		}

		return &Repository{repo: dbstorage.New(db)}, nil
	case "file":
		filename := cfg.DataSourceName
		if name, ok := os.LookupEnv("FILE_STORAGE_PATH"); ok {
			filename = name
		}

		return &Repository{repo: filestorage.New(filename)}, nil

	default:
		return &Repository{repo: mapstorage.New()}, nil
	}
}
