package repository

import (
	"fuegobyp-billing.com/internal/model"
)

type CustomerRepositoryInterface interface {
	Request(id string) (model.Customer, error)
}
