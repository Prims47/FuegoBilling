package exporter

type ExporterProviderInterface interface {
	Save(fileName string, data []byte) error
	CanSave(exporterProviderName string) bool
}
