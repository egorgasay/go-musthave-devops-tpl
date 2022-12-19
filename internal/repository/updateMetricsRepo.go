package repository

import (
	"errors"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)


func (ms MemStorageMethods) UpdateMetric(mt *Metrics) (err error) {
	var status sql.Result
	if mt.Type == "gauge" {
		queryUpdate :=
			`UPDATE metrics 
		SET value = ? 
		WHERE name = ?;`

		status, err = ms.DB.Exec(queryUpdate, mt.Value, mt.Name)
	} else if mt.Type == "counter" {
		queryIncrement :=
			`UPDATE metrics 
		SET value = ? + (SELECT value FROM metrics WHERE name = ?)
		WHERE name = ?;`

		status, err = ms.DB.Exec(queryIncrement, mt.Value, mt.Name, mt.Name)
	} else {
		return errors.New("метрики не существует")
	}

	if st, _ := status.RowsAffected(); st == 0 || err != nil {
		return errors.New("ошибка при обновлении "+err.Error())
	}
	
	return nil
}
