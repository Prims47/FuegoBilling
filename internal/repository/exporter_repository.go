package repository

import "github.com/prims47/FuegoBilling/internal/exporter"

type ExporterRepository struct {
	ExporterContext exporter.ExporterContextInterface
}

func NewExporterRepositoryRepository(exporterContext exporter.ExporterContextInterface) ExporterRepositoryInterface {
	return &ExporterRepository{ExporterContext: exporterContext}
}

func (e *ExporterRepository) Save(fileName string, exporterName string, data []byte) error {
	return e.ExporterContext.Save(fileName, exporterName, data)
}
