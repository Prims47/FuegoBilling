package model

type TVA struct {
	Pourcent float32
}

func (t *TVA) GetTVA(totalHT float32) float32 {
	return (totalHT * t.Pourcent) / 100
}
