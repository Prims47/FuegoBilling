package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat32ToString(t *testing.T) {
	t.Parallel()

	// Given
	expectedValue := "13.7"
	sut := FormatFloat{}

	// When

	formattedValue := sut.Float32ToString(float32(13.7))

	// Then

	assert.Equal(t, expectedValue, formattedValue)
}
