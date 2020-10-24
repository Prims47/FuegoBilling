package exporter

import (
	"errors"
)

type ExporterContext struct {
	providers []ExporterProviderInterface
}

func NewExporterProviderContext(providers []ExporterProviderInterface) ExporterContextInterface {
	return &ExporterContext{providers: providers}
}

func (e *ExporterContext) Save(exporterName string, data []byte) error {
	for _, provider := range e.providers {
		if provider.CanSave(exporterName) {
			provider.Save(data)
			return nil
		}
	}

	return errors.New("No provider found")
}
