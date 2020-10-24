package repository

import (
	"github.com/prims47/FuegoBilling/internal/model"
)

type CustomerRepositoryInterface interface {
	Request(id string) (model.Customer, error)
}
