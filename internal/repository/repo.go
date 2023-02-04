package repository

import (
	"database/sql"
	"devtool/internal/storage"
	dbstorage "devtool/internal/storage/db"
	filestorage "devtool/internal/storage/file"
	"errors"
	"os"
	"sync"
	"time"
)

type Config struct {
	DriverName     string
	DataSourceName string
	SaveAfter      time.Duration
	Restore        bool
}

type Repository struct {
	Mu            sync.RWMutex
	Store         map[string]float64
	BackupStorage storage.IBackupStorage
}

func New(cfg *Config) (*Repository, error) {
	if cfg == nil {
		panic("конфигурация задана некорректно")
	}

	ms := make(map[string]float64)
	repo := &Repository{
		Store: ms,
	}

	if cfg.Restore {
		defer repo.Restore()
	}

	switch cfg.DriverName {
	case "sqlite3":
		db, err := sql.Open(cfg.DriverName, cfg.DataSourceName)
		if err != nil {
			return nil, err
		}
		realDB := dbstorage.New(db)
		repo.BackupStorage = realDB

		return repo, nil
	case "file":
		filename := cfg.DataSourceName
		if name, ok := os.LookupEnv("FILE_STORAGE_PATH"); ok {
			filename = name
		}
		fileStorage := filestorage.New(filename)
		repo.BackupStorage = fileStorage

		return repo, nil
	default:
		return repo, errors.New("неподдерживаемый тип хранилища")
	}
}
