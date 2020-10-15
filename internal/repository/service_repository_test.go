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

func TestServiceRequest(t *testing.T) {
	t.Parallel()

	// Given

	mockCtrl := gomock.NewController(t)

	adapterMock := mock.NewMockServiceAdapterInterface(mockCtrl)

	testCases := []struct {
		testName        string
		params          string
		times           int
		adapterResponse adapter.ServiceAdapterResponse
		model           model.Service
		errAdapter      error
		errRepository   string
	}{
		{testName: "Given empty params", params: "", times: 1, adapterResponse: adapter.ServiceAdapterResponse{}, errRepository: "Please give a valid params"},
		{testName: "Given adapter return error", params: "id", times: 1, adapterResponse: adapter.ServiceAdapterResponse{}, errAdapter: errors.New("Error adapter"), errRepository: "Error adapter"},
		{testName: "Given adapter return without Detail",
			params:          "id",
			times:           1,
			adapterResponse: adapter.ServiceAdapterResponse{Detail: ""},
			errRepository:   "Error when parse Service. Data is not correct",
		},
		{testName: "Given adapter return without Quantity",
			params:          "id",
			times:           1,
			adapterResponse: adapter.ServiceAdapterResponse{Detail: "Prestation Pepito Fuego by P"},
			errRepository:   "Error when parse Service. Data is not correct",
		},
		{testName: "Given adapter return without Unit Price",
			params:          "id",
			times:           1,
			adapterResponse: adapter.ServiceAdapterResponse{Detail: "Prestation Pepito Fuego by P", Quantity: 1},
			errRepository:   "Error when parse Service. Data is not correct",
		},
		{testName: "Given adapter return without Tva",
			params:          "id",
			times:           1,
			adapterResponse: adapter.ServiceAdapterResponse{Detail: "Prestation Pepito Fuego by P", Quantity: 1, UnitPrice: 77047},
			errRepository:   "Error when parse Service. Data is not correct",
		},
		{testName: "Given adapter return with correct respobse",
			params:          "id",
			times:           1,
			adapterResponse: adapter.ServiceAdapterResponse{Detail: "Prestation Pepito Fuego by P", Quantity: 1, UnitPrice: 77047, TVA: adapter.TVAServiceResponse{Pourcent: 20}},
			model: model.Service{
				Detail:    "Prestation Pepito Fuego by P",
				Quantity:  1,
				UnitPrice: 77047,
				TVA:       model.TVA{Pourcent: 20},
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

			sut := NewServiceRepository(adapterMock)

			modelAccount, err := sut.Request(tc.params)

			assert.Equal(t, tc.model, modelAccount)

			if tc.errRepository != "" {
				assert.EqualError(t, err, tc.errRepository)
			}
		})

	}
}
