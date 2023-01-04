package filestorage

import (
	"bufio"
	"devtool/internal/globals"
	"devtool/internal/storage"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func (fs *FileStorage) UpdateOneMetric(mt *storage.Metrics) (count float64, err error) {
	if globals.Restore {
		err = fs.OpenRead()
	} else {
		err = fs.OpenWrite()
	}

	if err != nil {
		return 0, err
	}

	var lines = make([]string, 0)

	scanner := bufio.NewScanner(fs.File)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	flag := false
	for i, line := range lines {
		row := strings.Split(line, " ")
		if len(row) != 2 {
			continue
		}

		if row[0] == mt.ID {
			val, err := strconv.ParseFloat(row[1], 64)
			if err != nil {
				break
			}
			flag = true

			if mt.MType == "gauge" {
				lines[i] = fmt.Sprintf("%s %f", mt.ID, *mt.Value)
				count = *mt.Value + val
			} else if mt.MType == "counter" {
				lines[i] = fmt.Sprintf("%s %d", mt.ID, *mt.Delta+int64(val))
				count = float64(*mt.Delta + int64(val))
			}
			break
		}
	}

	if !flag {
		if mt.MType == "gauge" {
			lines = append(lines, fmt.Sprintf("%s %f", mt.ID, *mt.Value))
			count = *mt.Value
		} else if mt.MType == "counter" {
			lines = append(lines, fmt.Sprintf("%s %d", mt.ID, *mt.Delta))
			count = float64(*mt.Delta)
		}
	}
	log.Println(fs.Store[mt.ID], count)
	fs.Store[mt.ID] = count
	fs.Close()

	go func() {
		time.Sleep(globals.SaveAfter)

		fs.OpenWrite()
		defer fs.Close()
		output := []byte(strings.Join(lines, "\n"))

		writer := bufio.NewWriter(fs.File)
		writer.Write(output)
		writer.Flush()
	}()

	return count, nil
}

//
//func (fs *FileStorage) UpdateOneMetric2(mt *storage.Metrics) (count float64, err error) {
//	input, err := os.ReadFile(fs.Path)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	lines := strings.Split(string(input), "\n")
//
//	flag := false
//	for i, line := range lines {
//		if strings.Contains(line, mt.ID) {
//			flag = true
//			val, err := strconv.ParseFloat(strings.Split(line, " ")[1], 64)
//			if mt.MType == "gauge" {
//				lines[i] = fmt.Sprintf("%s %f", mt.ID, *mt.Value)
//				count = *mt.Value + val
//			} else if mt.MType == "counter" {
//				if err != nil {
//					continue
//				}
//				lines[i] = fmt.Sprintf("%s %d", mt.ID, *mt.Delta+int64(val))
//				count = float64(*mt.Delta + int64(val))
//			}
//		}
//	}
//
//	if !flag {
//		err = fs.Open()
//		if err != nil {
//			return 0, err
//		}
//		defer fs.Close()
//
//		writer := bufio.NewWriter(fs.File)
//		if mt.MType == "gauge" {
//			count = *mt.Value
//			writer.WriteString(fmt.Sprintf("%s %f\n", mt.ID, *mt.Value))
//		} else if mt.MType == "counter" {
//			count = float64(*mt.Delta)
//			writer.WriteString(fmt.Sprintf("%s %d\n", mt.ID, *mt.Delta))
//		}
//
//		err = writer.Flush()
//		if err != nil {
//			fs.Close()
//			return 0, err
//		}
//	} else {
//		output := strings.Join(lines, "\n")
//
//		f, err := os.OpenFile(fs.Path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer f.Close()
//
//		b := []byte(output)
//
//		scanner := bufio.NewScanner(f)
//		for scanner.Scan() {
//			b = append(b, scanner.Bytes()...)
//		}
//
//		os.WriteFile(fs.Path, b, 0644)
//	}
//
//	return count, nil
//}
