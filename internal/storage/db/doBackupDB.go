package dbstorage

import (
	"database/sql"
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

	for name, value := range ms {
		_, err = stmtIn.Exec(name, value)
		if err != nil {
			continue
		}

		_, err = stmtUp.Exec(value, name)
		if err != nil {
			continue
		}
	}

	return err
}
