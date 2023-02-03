package usecase

import (
	"devtool/internal/service"
	"errors"
)

type UseCase struct {
	service *service.Service
}

var NotFoundErr = errors.New("not found")

func New(service *service.Service) UseCase {
	return UseCase{service: service}
}
