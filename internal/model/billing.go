package model

import (
	"strings"

	"github.com/google/uuid"
)

type Billing struct{}

func generateUUID() string {
	id, _ := uuid.NewRandom()

	return id.String()[0:12]
}

func (b *Billing) GetBillingNumber() string {
	return "FR-" + strings.ToUpper(generateUUID())
}
