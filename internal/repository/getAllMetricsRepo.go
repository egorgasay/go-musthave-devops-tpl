package repository

func (ms MemStorage) GetAllMetrics() ([]*Metrics, error) {
	var mt []*Metrics
	query := "SELECT name, value FROM metrics"

	rows, err := ms.DB.Query(query)
	if err != nil {
		return mt, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		metric := &Metrics{MType: "Test"}

		err = rows.Scan(&metric.ID, &metric.Value)
		if err != nil {
			return nil, err
		}

		mt = append(mt, metric)
	}

	return mt, nil
}
