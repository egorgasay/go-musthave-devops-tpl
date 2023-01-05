package filestorage

//func (fs *FileStorage) GetOneMetric(name string) (float64, error) {
//	err := fs.OpenRead()
//	if err != nil {
//		return 0, err
//	}
//	defer fs.Close()
//
//	if val, ok := fs.Store[name]; ok {
//		return val, nil
//	}
//
//	scanner := bufio.NewScanner(fs.File)
//	for scanner.Scan() {
//		row := strings.Split(scanner.Text(), " ")
//		if len(row) != 2 {
//			return 0, errors.New("wrong file format")
//		}
//
//		metricName, value := row[0], row[1]
//		if name == metricName {
//			val, err := strconv.ParseFloat(value, 64)
//			if err != nil {
//				return 0, errors.New("wrong row format")
//			}
//
//			return val, nil
//		}
//	}
//
//	return 0, errors.New("not found")
//}
