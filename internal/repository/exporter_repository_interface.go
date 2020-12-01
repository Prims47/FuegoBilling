package repository

type ExporterRepositoryInterface interface {
	Save(fileName string, exporterName string, data []byte) error
}
