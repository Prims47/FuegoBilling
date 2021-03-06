package adapter

type AddressCustomerResponse struct {
	Street  string `json:"street"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type CompanyCustomerResponse struct {
	Siret string `json:"siret"`
	Tva   string `json:"tva"`
	Type  string `json:"type"`
}

type CustomerAdapterResponse struct {
	Name    string                  `json:"name"`
	Address AddressCustomerResponse `json:"address"`
	Company CompanyCustomerResponse `json:"company"`
}

type CustomerAdapterInterface interface {
	Request(id string) (CustomerAdapterResponse, error)
}
