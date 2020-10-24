package repository

import (
	"github.com/prims47/FuegoBilling/internal/model"
)

type ServiceRepositoryInterface interface {
	Request(id string) (model.Service, error)
}
