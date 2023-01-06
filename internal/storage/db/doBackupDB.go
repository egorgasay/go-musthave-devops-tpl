package dbstorage

import (
	"database/sql"
	"devtool/internal/storage"
	"log"
)

func (rs *RealStorage) DoBackup(ms map[string]float64) (err error) {
	var stmtIn, stmtUp *sql.Stmt

	stmtIn, err = rs.DB.Prepare(`
		INSERT OR IGNORE INTO metrics (name, value) VALUES (?, ?);
		`)
	if err != nil {
		return err
	}

	stmtUp, err = rs.DB.Prepare(`UPDATE metrics SET value = ? WHERE name = ?`)
	if err != nil {
		return err
	}

	storage.StorageRelevance.Mu.Lock()

	for name, value := range ms {
		if _, ok := storage.StorageRelevance.UpdateNeeded[name]; !ok {
			continue
		}

		delete(storage.StorageRelevance.UpdateNeeded, name)

		_, err = stmtIn.Exec(name, value)
		if err != nil {
			continue
		}

		_, err = stmtUp.Exec(value, name)
		if err != nil {
			continue
		}

		log.Println(name, "Saved")
	}

	storage.StorageRelevance.Mu.Unlock()

	return err
}
