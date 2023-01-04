package filestorage

import (
	"bufio"
	"devtool/internal/storage"
	"strconv"
	"strings"
)

func (fs *FileStorage) GetListOfMetrics() ([]*storage.Metrics, error) {
	var metrics []*storage.Metrics
	err := fs.OpenRead()
	if err != nil {
		return nil, err
	}
	defer fs.Close()

	scanner := bufio.NewScanner(fs.File)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		if len(row) != 2 {
			continue
			//return nil, errors.New("wrong file format")
		}

		metricName, value := row[0], row[1]
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			continue
			//return nil, errors.New("wrong row format")
		}
		metric := &storage.Metrics{
			ID:    metricName,
			Value: &val,
		}

		metrics = append(metrics, metric)
	}

	return metrics, nil
}
