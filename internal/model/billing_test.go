package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBillingNumber(t *testing.T) {
	t.Parallel()

	// Given

	sut := Billing{}

	// When

	billingNumber := sut.GetBillingNumber()

	// Then

	assert.NotEmpty(t, billingNumber)
	assert.Contains(t, billingNumber, "FR-")
}
