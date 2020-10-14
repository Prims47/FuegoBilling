package adapter

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type AccountAdapter struct {
}

func NewAccountJSONAdapter() AccountAdapterInterface {
	return &AccountAdapter{}
}

func (a *AccountAdapter) Request(id string) (AccountAdapterResponse, error) {
	if _, err := os.Stat(id); err != nil {
		return AccountAdapterResponse{}, errors.Errorf("Invalid account config path")
	}

	jsonFile, err := os.Open(id)

	defer jsonFile.Close()

	if err != nil {
		return AccountAdapterResponse{}, errors.Errorf("Invalid account config path")
	}

	bytesValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return AccountAdapterResponse{}, errors.Errorf("Impossible to read JSON file")
	}

	account := AccountAdapterResponse{}

	err = json.Unmarshal(bytesValue, &account)

	if err != nil {
		return AccountAdapterResponse{}, errors.Errorf("Impossible to map JSON file")
	}

	if (AccountAdapterResponse{}) == account {
		return AccountAdapterResponse{}, errors.Errorf("Impossible to Unmarshal JSON file")
	}

	if (AccountAdapterResponse{}.Address) == account.Address {
		return AccountAdapterResponse{}, errors.Errorf("Impossible to Unmarshal Address JSON file")
	}

	if (AccountAdapterResponse{}.Company) == account.Company {
		return AccountAdapterResponse{}, errors.Errorf("Impossible to Unmarshal Company JSON file")
	}

	if (AccountAdapterResponse{}.Name) == account.Name &&
		(AccountAdapterResponse{}.FirstName) == account.FirstName &&
		(AccountAdapterResponse{}.LastName) == account.LastName &&
		(AccountAdapterResponse{}.Mail) == account.Mail {
		return AccountAdapterResponse{}, errors.Errorf("Impossible to Unmarshal Info JSON file")
	}

	return account, nil
}
