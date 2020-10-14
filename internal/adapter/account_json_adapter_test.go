package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountRequest(t *testing.T) {
	t.Parallel()

	// Given

	expectedResponse := AccountAdapterResponse{
		Name:      "Fuego by P",
		FirstName: "Ilan",
		LastName:  "Zerath",
		Mail:      "pepito@fuegobyp.com",
		Address: AddressAccountResponse{
			City:    "Paris",
			Country: "France",
			Street:  "770 rue de Fuego",
			ZipCode: "75006",
		},
		Company: CompanyAccountResponse{
			Siret:   "88462068300018",
			Tva:     "FR77049322770",
			Capital: 100.00,
			RCS:     "Nanterre B",
			NAF:     "6201Z",
			Type:    "SARL",
		},
	}

	sut := NewAccountJSONAdapter()

	testCases := []struct {
		name     string
		id       string
		errorMsg string
		response AccountAdapterResponse
	}{
		{name: "Given No ID", id: "", errorMsg: "Invalid account config path"},
		{name: "Given not existed ID", id: "toto", errorMsg: "Invalid account config path"},
		{name: "Given empty JSON file", id: "../../tests/config/account_empty_test.json", errorMsg: "Impossible to map JSON file"},
		{name: "Given empty CONTENT JSON file", id: "../../tests/config/account_empty_content_test.json", errorMsg: "Impossible to Unmarshal JSON file"},
		{name: "Given uncompleted address JSON file", id: "../../tests/config/account_uncompleted_address_content.json", errorMsg: "Impossible to Unmarshal Address JSON file"},
		{name: "Given uncompleted company JSON file", id: "../../tests/config/account_uncompleted_company_content.json", errorMsg: "Impossible to Unmarshal Company JSON file"},
		{name: "Given uncompleted info JSON file", id: "../../tests/config/account_uncompleted_info.json", errorMsg: "Impossible to Unmarshal Info JSON file"},
		{name: "Given completed JSON file", id: "../../tests/config/account_completed_content.json", errorMsg: "", response: expectedResponse},
	}

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			response, err := sut.Request(tc.id)

			if tc.errorMsg != "" {
				assert.Equal(t, tc.errorMsg, err.Error())
			}

			assert.Equal(t, tc.response, response)
		})
	}
}
