package service

import (
	repo "devtool/internal/repository"
)

type Service struct {
	DB *repo.MemStorage
}

// ошибка из-за того что нул передаем на место интерфейса
// мб создать еще одну структуру?...
// type IServiceFunc interface {
// 	UpdateMetric(*Metrics) error
// 	GetMetric(string) (float64, error)
// 	GetAllMetrics([]Metrics) error
// }


func NewService(cfg repo.Config) (*Service, error) {
	ms, err := repo.NewMemStorage(&cfg)
	if err != nil {
		return nil, err
	}

	return &Service{DB: ms}, nil
}

type Metrics repo.Metrics

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
