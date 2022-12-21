package repository

import (
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

func (ms MemStorage) GetMetric(name string) (float64, error) {
	query := "SELECT value FROM metrics WHERE name = ?;"
	row := ms.DB.QueryRow(query, name)

	var val float64
	if err := row.Scan(&val); err != nil {
		return 0, errors.New("значение не установлено")
	}

	return val, nil
}
