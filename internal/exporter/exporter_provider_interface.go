package exporter

type ExporterProviderInterface interface {
	Save(data []byte) error
	CanSave(exporterProviderName string) bool
}
