package adapter

type TVAServiceResponse struct {
	Pourcent float32 `json:"pourcent"`
}

type ServiceAdapterResponse struct {
	Detail    string             `json:"detail"`
	Quantity  float32            `json:"quantity"`
	UnitPrice float32            `json:"unitPrice"`
	TVA       TVAServiceResponse `json:"tva"`
}

type ServiceAdapterInterface interface {
	Request(id string) (ServiceAdapterResponse, error)
}
