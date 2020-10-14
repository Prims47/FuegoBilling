package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToStringFrenchFormat(t *testing.T) {
	t.Parallel()

	// Given
	expectedValue := "13 760"
	sut := FormatInt{}

	// When

	formattedValue := sut.IntToStringFrenchFormat(13760)

	// Then

	assert.Equal(t, expectedValue, formattedValue)
}
