package exporter

type ExporterContextInterface interface {
	Save(exporterName string, data []byte) error
}
