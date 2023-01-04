package filestorage

import (
	"os"
	"sync"
)

type FileStorage struct {
	Path  string
	File  *os.File
	Store map[string]float64
	Mu    sync.Mutex
}

func New(path string) *FileStorage {
	return &FileStorage{
		Path:  path,
		Store: make(map[string]float64),
	}
}

func (fs *FileStorage) OpenRead() error {
	fs.Mu.Lock()
	file, err := os.OpenFile(fs.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	fs.File = file
	return nil
}

func (fs *FileStorage) OpenWrite() error {
	fs.Mu.Lock()
	file, err := os.OpenFile(fs.Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	fs.File = file
	return nil
}

func (fs *FileStorage) Close() error {
	fs.Mu.Unlock()
	return fs.File.Close()
}
