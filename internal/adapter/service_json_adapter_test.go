package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceRequest(t *testing.T) {
	t.Parallel()

	// Given

	expectedResponse := ServiceAdapterResponse{
		Detail:    "Prestation Pepito Fuego by P",
		Quantity:  1,
		UnitPrice: 77047,
		TVA: TVAServiceResponse{
			Pourcent: 77047,
		},
	}

	sut := NewServiceJSONAdapter()

	testCases := []struct {
		name     string
		id       string
		errorMsg string
		response ServiceAdapterResponse
	}{
		{name: "Given No ID", id: "", errorMsg: "Invalid service config path"},
		{name: "Given not existed ID", id: "toto", errorMsg: "Invalid service config path"},
		{name: "Given empty JSON file", id: "../../tests/config/service_empty_test.json", errorMsg: "Impossible to map JSON file"},
		{name: "Given empty CONTENT JSON file", id: "../../tests/config/service_empty_content_test.json", errorMsg: "Impossible to Unmarshal JSON file"},
		{name: "Given uncompleted tva JSON file", id: "../../tests/config/service_uncompleted_tva_content.json", errorMsg: "Impossible to Unmarshal TVA JSON file"},
		{name: "Given uncompleted info JSON file", id: "../../tests/config/service_uncompleted_info.json", errorMsg: "Impossible to Unmarshal JSON file"},
		{name: "Given completed JSON file", id: "../../tests/config/service_completed_content.json", errorMsg: "", response: expectedResponse},
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
