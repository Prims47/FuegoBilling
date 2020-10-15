package adapter

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type ServiceAdapter struct {
}

func NewServiceJSONAdapter() ServiceAdapterInterface {
	return &ServiceAdapter{}
}

func (s *ServiceAdapter) Request(id string) (ServiceAdapterResponse, error) {
	if _, err := os.Stat(id); err != nil {
		return ServiceAdapterResponse{}, errors.Errorf("Invalid service config path")
	}

	jsonFile, err := os.Open(id)

	defer jsonFile.Close()

	if err != nil {
		return ServiceAdapterResponse{}, errors.Errorf("Invalid service config path")
	}

	bytesValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return ServiceAdapterResponse{}, errors.Errorf("Impossible to read JSON file")
	}

	service := ServiceAdapterResponse{}

	err = json.Unmarshal(bytesValue, &service)

	if err != nil {
		return ServiceAdapterResponse{}, errors.Errorf("Impossible to map JSON file")
	}

	if (ServiceAdapterResponse{}) == service ||
		(ServiceAdapterResponse{}.Detail) == service.Detail ||
		(ServiceAdapterResponse{}.Quantity) == service.Quantity ||
		(ServiceAdapterResponse{}.UnitPrice) == service.UnitPrice {
		return ServiceAdapterResponse{}, errors.Errorf("Impossible to Unmarshal JSON file")
	}

	if (ServiceAdapterResponse{}.TVA) == service.TVA {
		return ServiceAdapterResponse{}, errors.Errorf("Impossible to Unmarshal TVA JSON file")
	}

	return service, nil
}
