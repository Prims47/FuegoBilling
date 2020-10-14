package adapter

type AddressAccountResponse struct {
	Street  string `json:"street"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type CompanyAccountResponse struct {
	Siret   string  `json:"siret"`
	Tva     string  `json:"tva"`
	Capital float32 `json:"capital"`
	RCS     string  `json:"rcs"`
	NAF     string  `json:"naf"`
	Type    string  `json:"type"`
}

type AccountAdapterResponse struct {
	Name      string                 `json:"name"`
	FirstName string                 `json:"firstname"`
	LastName  string                 `json:"lastname"`
	Address   AddressAccountResponse `json:"address"`
	Company   CompanyAccountResponse `json:"company"`
	Mail      string                 `json:"mail"`
}

type AccountAdapterInterface interface {
	Request(id string) (AccountAdapterResponse, error)
}
