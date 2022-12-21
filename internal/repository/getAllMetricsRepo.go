package repository

func (ms MemStorage) GetAllMetrics() ([]*Metrics, error) {
	var mt []*Metrics
	query := "SELECT name, value FROM metrics"
	rows, err := ms.DB.Query(query)
	if err != nil {
		return mt, err
	}

	for rows.Next() {
		metric := &Metrics{Type: "Test"}
		rows.Scan(&metric.Name, &metric.Value)
		mt = append(mt, metric)
	}
	return mt, nil
}
