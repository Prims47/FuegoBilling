package repository

import (
	"github.com/prims47/FuegoBilling/internal/model"
)

type AccountRepositoryInterface interface {
	Request(id string) (model.Account, error)
}
