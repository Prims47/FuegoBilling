package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTVA(t *testing.T) {
	t.Parallel()

	// Given

	expectedTVA := float32(2750)
	sut := TVA{Pourcent: 20}

	// When

	tva := sut.GetTVA(13750)

	// Then

	assert.Equal(t, expectedTVA, tva)

}
