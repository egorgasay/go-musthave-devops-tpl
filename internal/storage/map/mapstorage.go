package mapstorage

import "sync"

type MapStorage struct {
	mu    sync.RWMutex
	store map[string]float64
}

func New() *MapStorage {
	ms := make(map[string]float64)
	return &MapStorage{
		store: ms,
	}
}
