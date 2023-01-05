package filestorage

import (
	"bufio"
	"fmt"
)

func (fs *FileStorage) DoBackup(ms map[string]float64) error {
	fs.OpenWrite()
	defer fs.Close()

	writer := bufio.NewWriter(fs.File)
	for name, value := range ms {
		_, err := writer.Write([]byte(fmt.Sprintln(name, value)))
		if err != nil {
			return err
		}
	}

	err := writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
