package repository

import (
	"errors"
	"testing"

	"fuegobyp-billing.com/internal/adapter"
	"fuegobyp-billing.com/internal/model"
	mock "fuegobyp-billing.com/internal/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAccountRequest(t *testing.T) {
	t.Parallel()

	// Given

	mockCtrl := gomock.NewController(t)

	adapterMock := mock.NewMockAccountAdapterInterface(mockCtrl)

	testCases := []struct {
		testName        string
		params          string
		times           int
		adapterResponse adapter.AccountAdapterResponse
		model           model.Account
		errAdapter      error
		errRepository   string
	}{
		{testName: "Given empty params", params: "", times: 1, adapterResponse: adapter.AccountAdapterResponse{}, errRepository: "Please give a valid params"},
		{testName: "Given adapter return error", params: "id", times: 1, adapterResponse: adapter.AccountAdapterResponse{}, errAdapter: errors.New("Error adapter"), errRepository: "Error adapter"},
		{testName: "Given adapter return without name",
			params:          "id",
			times:           1,
			adapterResponse: adapter.AccountAdapterResponse{Name: ""},
			errRepository:   "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without firstname",
			params:          "id",
			times:           1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito", FirstName: ""},
			errRepository:   "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without lastname",
			params:          "id",
			times:           1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito", FirstName: "Ilan", LastName: ""},
			errRepository:   "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without mail",
			params:          "id",
			times:           1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito", FirstName: "Ilan", LastName: "Zerath", Mail: ""},
			errRepository:   "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without Address Street",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street: "",
				},
			},
			errRepository: "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without Address ZipCode",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street:  "770 rue du Fuego",
					ZipCode: "",
				},
			},
			errRepository: "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without Address City",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "",
				},
			},
			errRepository: "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without Address Country",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "",
				},
			},
			errRepository: "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without Company Siret",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "France",
				},
				Company: adapter.CompanyAccountResponse{
					Siret: "",
				},
			},
			errRepository: "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without Company Tva",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "France",
				},
				Company: adapter.CompanyAccountResponse{
					Siret: "11212",
					Tva:   "",
				},
			},
			errRepository: "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without Company Capital",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "France",
				},
				Company: adapter.CompanyAccountResponse{
					Siret: "11212",
					Tva:   "21212",
				},
			},
			errRepository: "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without Company RCS",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "France",
				},
				Company: adapter.CompanyAccountResponse{
					Siret:   "11212",
					Tva:     "21212",
					Capital: 100.47,
					RCS:     "",
				},
			},
			errRepository: "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without Company NAF",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "France",
				},
				Company: adapter.CompanyAccountResponse{
					Siret:   "11212",
					Tva:     "21212",
					Capital: 100.47,
					RCS:     "Paris",
					NAF:     "",
				},
			},
			errRepository: "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter return without Company Type",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{Name: "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "France",
				},
				Company: adapter.CompanyAccountResponse{
					Siret:   "11212",
					Tva:     "21212",
					Capital: 100.47,
					RCS:     "Paris",
					NAF:     "NAF",
					Type:    "",
				},
			},
			errRepository: "Error when parse Account. Data is not correct",
		},
		{testName: "Given adapter correct response",
			params: "id",
			times:  1,
			adapterResponse: adapter.AccountAdapterResponse{
				Name:      "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: adapter.AddressAccountResponse{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "France",
				},
				Company: adapter.CompanyAccountResponse{
					Siret:   "11212",
					Tva:     "21212",
					Capital: 100.47,
					RCS:     "Paris",
					NAF:     "NAF",
					Type:    "SARL",
				},
			},
			model: model.Account{
				Name:      "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: model.Address{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "France",
				},
				Company: model.Company{
					Siret:   "11212",
					Tva:     "21212",
					Capital: 100.47,
					RCS:     "Paris",
					NAF:     "NAF",
					Type:    "SARL",
				},
			},
		},
	}

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			adapterMock.EXPECT().
				Request(gomock.Eq(tc.params)).
				Times(tc.times).
				Return(tc.adapterResponse, tc.errAdapter)

			sut := NewAccountRepository(adapterMock)

			modelAccount, err := sut.Request(tc.params)

			assert.Equal(t, tc.model, modelAccount)

			if tc.errRepository != "" {
				assert.EqualError(t, err, tc.errRepository)
			}
		})

	}
}
