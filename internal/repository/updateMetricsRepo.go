package repository

import (
	"errors"

	_ "github.com/mattn/go-sqlite3"
)


func (ms MemStorageMethods) UpdateMetric(mt *Metrics) (err error) {
	if mt.Type == "gauge" {
		queryUpdate :=
			`INSERT OR IGNORE INTO metrics(name, value) VALUES(?, 0);
		UPDATE metrics 
		SET value = ? 
		WHERE name = ?;`

		_, err = ms.DB.Exec(queryUpdate, mt.Name, mt.Value, mt.Name)
	} else if mt.Type == "counter" {
		queryIncrement :=
			`INSERT OR IGNORE INTO metrics(name, value) VALUES(?, 0);
		UPDATE metrics 
		SET value = ? + (SELECT value FROM metrics WHERE name = ?)
		WHERE name = ?;`

		_, err = ms.DB.Exec(queryIncrement, mt.Name, mt.Value, mt.Name, mt.Name)
	} else {
		return errors.New("метрики не существует")
	}
	
	return err
}




// if st, _ := status.RowsAffected(); st == 0 {
// 	queryCreate :=
// 		`UPDATE metrics 
// 	SET value = ? 
// 	WHERE name = ?;`

// 	status, err = ms.DB.Exec(queryCreate, mt.Value, mt.Name)
// }