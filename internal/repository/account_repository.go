package repository

import (
	"errors"

	"github.com/prims47/FuegoBilling/internal/adapter"
	"github.com/prims47/FuegoBilling/internal/model"
)

type AccountRepository struct {
	Adapter adapter.AccountAdapterInterface
}

func NewAccountRepository(adapter adapter.AccountAdapterInterface) AccountRepositoryInterface {
	return &AccountRepository{Adapter: adapter}
}

func (a *AccountRepository) Request(id string) (model.Account, error) {
	if id == "" {
		return model.Account{}, errors.New("Please give a valid params")
	}

	accountAdapterResponse, err := a.Adapter.Request(id)

	if err != nil {
		return model.Account{}, err
	}

	return parseAccountAdapterResponse(accountAdapterResponse)
}

func parseAccountAdapterResponse(response adapter.AccountAdapterResponse) (model.Account, error) {
	if response.Name == "" ||
		response.FirstName == "" ||
		response.LastName == "" ||
		response.Mail == "" ||
		response.Address.Street == "" ||
		response.Address.ZipCode == "" ||
		response.Address.City == "" ||
		response.Address.Country == "" ||
		response.Company.Siret == "" ||
		response.Company.Tva == "" ||
		response.Company.Capital < 1 ||
		response.Company.RCS == "" ||
		response.Company.NAF == "" ||
		response.Company.Type == "" {
		return model.Account{}, errors.New("Error when parse Account. Data is not correct")
	}

	return mapAccountAdapterResponseToModel(response), nil
}

func mapAccountAdapterResponseToModel(response adapter.AccountAdapterResponse) model.Account {
	return model.Account{
		Name:      response.Name,
		FirstName: response.FirstName,
		LastName:  response.LastName,
		Mail:      response.Mail,
		Address: model.Address{
			Street:  response.Address.Street,
			City:    response.Address.City,
			Country: response.Address.Country,
			ZipCode: response.Address.ZipCode,
		},
		Company: model.Company{
			Siret:   response.Company.Siret,
			Tva:     response.Company.Tva,
			Capital: response.Company.Capital,
			NAF:     response.Company.NAF,
			RCS:     response.Company.RCS,
			Type:    response.Company.Type,
		},
	}
}
