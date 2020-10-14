package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceTotalHT(t *testing.T) {
	t.Parallel()

	// Given

	expectedTotalHT := float32(13260)
	tva := TVA{Pourcent: 20}
	sut := Service{Detail: "Prestation de développement Bivwak BNP Paribas", Quantity: 20, UnitPrice: 663, TVA: tva}

	// When

	totalHT := sut.GetTotalHT()

	// Then

	assert.Equal(t, expectedTotalHT, totalHT)

}

func TestServiceTotalTTC(t *testing.T) {
	t.Parallel()

	// Given

	expectedTotalTTC := float32(15912)
	tva := TVA{Pourcent: 20}
	sut := Service{Detail: "Prestation de développement Bivwak BNP Paribas", Quantity: 20, UnitPrice: 663, TVA: tva}

	// When

	totalTTC := sut.GetTotalTTC()

	// Then

	assert.Equal(t, expectedTotalTTC, totalTTC)
}
