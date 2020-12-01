package providersExporter

import (
	"github.com/prims47/FuegoBilling/internal/exporter"
)

type RegisterExporter struct{}

func NewRegisterExporter() *RegisterExporter {
	return &RegisterExporter{}
}

func (p *RegisterExporter) Register() []exporter.ExporterProviderInterface {
	return []exporter.ExporterProviderInterface{
		&LocalExporter{},
		&AWSExporter{},
		&GoogleDriveExporter{},
	}
}
