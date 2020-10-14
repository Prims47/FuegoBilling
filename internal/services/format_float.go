package services

import "fmt"

type FormatFloatInterface interface {
	Float32ToString(value float32) string
}

type FormatFloat struct{}

func (f *FormatFloat) Float32ToString(value float32) string {
	return fmt.Sprintf("%g", value)
}
