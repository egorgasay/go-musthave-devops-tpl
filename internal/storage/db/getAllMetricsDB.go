package dbstorage

import "devtool/internal/storage"

func (rs *RealStorage) GetAllMetrics() ([]*storage.Metrics, error) {
	var mt []*storage.Metrics
	query := "SELECT name, value FROM metrics"

	rows, err := rs.DB.Query(query)
	if err != nil {
		return mt, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		metric := &storage.Metrics{MType: "Test"}

		err = rows.Scan(&metric.ID, &metric.Value)
		if err != nil {
			return nil, err
		}

		mt = append(mt, metric)
	}

	return mt, nil
}
