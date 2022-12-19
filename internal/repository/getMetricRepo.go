package repository

import (
	_ "github.com/mattn/go-sqlite3"
)

func (ms MemStorage) GetMetric(name string) (float64, error) {
	query := "SELECT value FROM metrics WHERE name = ?;"
	row := ms.DB.QueryRow(query, name)

	var val float64
	if err := row.Scan(&val); err != nil {
		return 0, err
	}

	return val, nil
}
