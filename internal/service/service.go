package service

//go:generate mockgen -source=service.go -destination=mocks/mock.go

import (
	repo "devtool/internal/repository"
)

type Service struct {
	UpdateMetric
	GetMetric
	GetAllMetrics
}

func NewService(db *repo.MemStorage) *Service {
	return &Service{UpdateMetric: db,
		GetMetric:     db,
		GetAllMetrics: db}
}

type Metrics repo.Metrics

//type IService interface {
//	UpdateMetric(*repo.Metrics) error
//	GetMetric(string) (float64, error)
//	GetAllMetrics([]repo.Metrics) error
//}

type UpdateMetric interface {
	UpdateMetric(*repo.Metrics) error
}

type GetMetric interface {
	GetMetric(string) (float64, error)
}

type GetAllMetrics interface {
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
