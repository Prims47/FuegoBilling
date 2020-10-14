package model

type Service struct {
	Detail    string
	Quantity  float32
	UnitPrice float32
	TVA       TVA
}

func (s *Service) GetTotalHT() float32 {
	return s.Quantity * s.UnitPrice
}

func (s *Service) GetTotalTTC() float32 {
	return s.GetTotalHT() + s.TVA.GetTVA(s.GetTotalHT())
}
