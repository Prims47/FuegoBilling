package repository

import (
	"fuegobyp-billing.com/internal/model"
)

type ServiceRepositoryInterface interface {
	Request(id string) (model.Service, error)
}
