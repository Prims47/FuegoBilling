package repository

import (
	"errors"

	"fuegobyp-billing.com/internal/adapter"
	"fuegobyp-billing.com/internal/model"
)

type CustomerRepository struct {
	Adapter adapter.CustomerAdapterInterface
}

func NewCustomerRepository(adapter adapter.CustomerAdapterInterface) CustomerRepositoryInterface {
	return &CustomerRepository{Adapter: adapter}
}

func (c *CustomerRepository) Request(id string) (model.Customer, error) {
	if id == "" {
		return model.Customer{}, errors.New("Please give a valid params")
	}

	customerAdapterResponse, err := c.Adapter.Request(id)

	if err != nil {
		return model.Customer{}, err
	}

	return parseCustomerAdapterResponse(customerAdapterResponse)
}

func parseCustomerAdapterResponse(response adapter.CustomerAdapterResponse) (model.Customer, error) {
	if response.Name == "" ||
		response.Address.Street == "" ||
		response.Address.ZipCode == "" ||
		response.Address.City == "" ||
		response.Address.Country == "" ||
		response.Company.Siret == "" ||
		response.Company.Tva == "" ||
		response.Company.Type == "" {
		return model.Customer{}, errors.New("Error when parse Customer. Data is not correct")
	}

	return mapCustomerAdapterResponseToModel(response), nil
}

func mapCustomerAdapterResponseToModel(response adapter.CustomerAdapterResponse) model.Customer {
	return model.Customer{
		Name: response.Name,
		Address: model.Address{
			Street:  response.Address.Street,
			City:    response.Address.City,
			Country: response.Address.Country,
			ZipCode: response.Address.ZipCode,
		},
		Company: model.Company{
			Siret:   response.Company.Siret,
			Tva:     response.Company.Tva,
			Capital: 0,
			NAF:     "",
			RCS:     "",
			Type:    response.Company.Type,
		},
	}
}
