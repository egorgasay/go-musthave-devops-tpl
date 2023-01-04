package dbstorage

import (
	"database/sql"
)

type IRealStorage interface {
	Ping() error
	Close() error
	Query(string, ...any) (*sql.Rows, error)
	Exec(string, ...any) (sql.Result, error)
	QueryRow(string, ...any) *sql.Row
}

type RealStorage struct {
	DB IRealStorage
}

func New(db *sql.DB) *RealStorage {
	return &RealStorage{
		DB: db,
	}
}
