package service

import (
	repo "devtool/internal/repository"
)

type Service struct {
	IService
}

func NewService(ms repo.IMemStorage) *Service {
	return &Service{IService: repo.NewMemStorageMethods(ms)}
}

type Metrics repo.Metrics

type IService interface {
	UpdateMetric(*repo.Metrics) error
	GetMetric(string) (float64, error)
	GetAllMetrics([]repo.Metrics) error
}

// type Storage struct {
// 	DB IStorage
// }

// type UpdateMetric interface {
// 	UpdateMetric(Metrics) error
// }

// type GetMetrics interface {
// 	GetMetric(string) (float64, error)
// 	GetAllMetrics([]Metrics) error
// }

// type Service struct {
// 	UpdateMetric
// 	GetMetrics
// }

// ошибка из-за того что нул передаем на место интерфейса
// type IServiceFunc interface {
// 	UpdateMetric(*Metrics) error
// 	GetMetric(string) (float64, error)
// 	GetAllMetrics([]Metrics) error
// }
