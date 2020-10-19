package repository

import (
	"fuegobyp-billing.com/internal/model"
)

type AccountRepositoryInterface interface {
	Request(id string) (model.Account, error)
}
