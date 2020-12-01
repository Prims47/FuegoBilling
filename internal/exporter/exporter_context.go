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

func (e *ExporterContext) Save(fileName string, exporterName string, data []byte) error {
	for _, provider := range e.providers {
		if provider.CanSave(exporterName) {
			return provider.Save(fileName, data)
		}
	}

	return errors.New("No provider found")
}
