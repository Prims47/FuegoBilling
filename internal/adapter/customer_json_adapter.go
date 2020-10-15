package adapter

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type CustomerAdapter struct {
}

func NewCustomerJSONAdapter() CustomerAdapterInterface {
	return &CustomerAdapter{}
}

func (c *CustomerAdapter) Request(id string) (CustomerAdapterResponse, error) {
	if _, err := os.Stat(id); err != nil {
		return CustomerAdapterResponse{}, errors.Errorf("Invalid customer config path")
	}

	jsonFile, err := os.Open(id)

	defer jsonFile.Close()

	if err != nil {
		return CustomerAdapterResponse{}, errors.Errorf("Invalid customer config path")
	}

	bytesValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return CustomerAdapterResponse{}, errors.Errorf("Impossible to read JSON file")
	}

	customer := CustomerAdapterResponse{}

	err = json.Unmarshal(bytesValue, &customer)

	if err != nil {
		return CustomerAdapterResponse{}, errors.Errorf("Impossible to map JSON file")
	}

	if (CustomerAdapterResponse{}) == customer {
		return CustomerAdapterResponse{}, errors.Errorf("Impossible to Unmarshal JSON file")
	}

	if (CustomerAdapterResponse{}.Address) == customer.Address {
		return CustomerAdapterResponse{}, errors.Errorf("Impossible to Unmarshal Address JSON file")
	}

	if (CustomerAdapterResponse{}.Company) == customer.Company {
		return CustomerAdapterResponse{}, errors.Errorf("Impossible to Unmarshal Company JSON file")
	}

	if (CustomerAdapterResponse{}.Name) == customer.Name {
		return CustomerAdapterResponse{}, errors.Errorf("Impossible to Unmarshal Info JSON file")
	}

	return customer, nil
}
