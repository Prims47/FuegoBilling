package repository

import (
	"errors"

	"fuegobyp-billing.com/internal/adapter"
	"fuegobyp-billing.com/internal/model"
)

type ServiceRepositoryInterface interface {
	Request(id string) (model.Service, error)
}

type ServiceRepository struct {
	Adapter adapter.ServiceAdapterInterface
}

func NewServiceRepository(adapter adapter.ServiceAdapterInterface) ServiceRepositoryInterface {
	return &ServiceRepository{Adapter: adapter}
}

func (s *ServiceRepository) Request(id string) (model.Service, error) {
	if id == "" {
		return model.Service{}, errors.New("Please give a valid params")
	}

	serviceAdapterResponse, err := s.Adapter.Request(id)

	if err != nil {
		return model.Service{}, err
	}

	return parseServiceAdapterResponse(serviceAdapterResponse)
}

func parseServiceAdapterResponse(response adapter.ServiceAdapterResponse) (model.Service, error) {
	if response.Detail == "" ||
		response.Quantity < 0.5 ||
		response.UnitPrice < 1 ||
		response.TVA.Pourcent < 1 {
		return model.Service{}, errors.New("Error when parse Service. Data is not correct")
	}

	return mapServiceAdapterResponseToModel(response), nil
}

func mapServiceAdapterResponseToModel(response adapter.ServiceAdapterResponse) model.Service {
	return model.Service{
		Detail:    response.Detail,
		Quantity:  response.Quantity,
		UnitPrice: response.UnitPrice,
		TVA: model.TVA{
			Pourcent: response.TVA.Pourcent,
		},
	}
}
