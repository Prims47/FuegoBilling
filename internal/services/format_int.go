package services

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type FormatIntInterface interface {
	IntToStringFrenchFormat(value int) string
}

type FormatInt struct{}

func (f *FormatInt) IntToStringFrenchFormat(value int) string {
	p := message.NewPrinter(language.French)

	return p.Sprintf("%d", value)
}
