package providersExporter

import (
	"io/ioutil"
	"os"
)

type LocalExporter struct{}

func (l *LocalExporter) Save(fileName string, data []byte) error {
	pdfPath := "generated-pdf"

	if _, err := os.Stat(pdfPath); err != nil {
		os.Mkdir(pdfPath, os.ModePerm)
	}

	return ioutil.WriteFile(pdfPath+"/"+fileName, data, 0644)
}

func (l *LocalExporter) CanSave(exporterProviderName string) bool {
	if "" == exporterProviderName || "local" == exporterProviderName {
		return true
	}

	return false
}
