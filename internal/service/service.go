package service

//go:generate mockgen -source=service.go -destination=mocks/mock.go

import (
	repo "devtool/internal/repository"
)

// type Service struct {
// 	UpdateMetric
// 	GetMetric
// 	GetAllMetrics
// }

type IService interface {
	UpdateMetric(*repo.Metrics) (float64, error)
	GetMetric(string) (float64, error)
	GetAllMetrics() ([]*repo.Metrics, error)
}

type Service struct {
	DB IService
}

func NewService(db *repo.MemStorage) *Service {
	return &Service{DB: db}
}

type Metrics repo.Metrics

// type UpdateMetric interface {
// 	UpdateMetric(*repo.Metrics) error
// }

// type GetMetric interface {
// 	GetMetric(string) (float64, error)
// }

// type GetAllMetrics interface {
// 	GetAllMetrics([]repo.Metrics) error
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
