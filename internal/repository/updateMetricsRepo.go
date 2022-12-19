package repository

import (
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

func (ms MemStorage) UpdateMetric(mt *Metrics) (err error) {
	if mt.Type == "gauge" {
		queryUpdate :=
			`UPDATE metrics 
		SET value = ? 
		WHERE name = ?;`
		_, err = ms.DB.Exec(queryUpdate, mt.Value, mt.Name)

		return
	} else if mt.Type == "counter" {
		queryIncrement :=
			`UPDATE metrics 
		SET value = ? + (SELECT value FROM metrics WHERE name = ?)
		WHERE name = ?;`
		_, err = ms.DB.Exec(queryIncrement, mt.Value, mt.Name, mt.Name)

		return
	}

	return errors.New("метрики не существует")
}
