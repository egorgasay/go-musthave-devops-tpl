package repository

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

func (ms MemStorage) UpdateMetric(mt *Metrics) (count float64, err error) {
	var row *sql.Row

	if mt.MType == "gauge" {
		queryUpdate :=
			`INSERT OR IGNORE INTO metrics(name, value) VALUES(?, 0);
		UPDATE metrics 
		SET value = ? 
		WHERE name = ?;
		`

		_, err = ms.DB.Exec(queryUpdate, mt.ID, mt.Value, mt.ID)
	} else if mt.MType == "counter" {
		queryIncrement :=
			`INSERT OR IGNORE INTO metrics(name, value) VALUES(?, 0);
		UPDATE metrics 
		SET value = ? + (SELECT value FROM metrics WHERE name = ?)
		WHERE name = ?;`

		_, err = ms.DB.Exec(queryIncrement, mt.ID, mt.Delta, mt.ID, mt.ID)
	} else {
		return 0, errors.New("тип не определен")
	}

	if err != nil {
		return 0, err
	}

	queryGetValue := `SELECT value FROM metrics WHERE name = ?;`
	row = ms.DB.QueryRow(queryGetValue, mt.ID)

	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
