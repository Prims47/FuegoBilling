package exporter

type ExporterContextInterface interface {
	Save(fileName string, exporterName string, data []byte) error
}
