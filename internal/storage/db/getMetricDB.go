package dbstorage

//func (rs *RealStorage) GetOneMetric(name string) (float64, error) {
//	query := "SELECT value FROM metrics WHERE name = ?;"
//	row := rs.DB.QueryRow(query, name)
//
//	var val float64
//	if err := row.Scan(&val); err != nil {
//		return 0, errors.New("значение не установлено")
//	}
//
//	return val, nil
//}
