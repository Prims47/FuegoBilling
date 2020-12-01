package providersExporter

import (
	"testing"

	"github.com/prims47/FuegoBilling/internal/exporter"
	"github.com/stretchr/testify/assert"
)

func TestRegisterExporter(t *testing.T) {
	t.Parallel()

	// Given

	expectedExporters := []exporter.ExporterProviderInterface{
		&LocalExporter{},
		&AWSExporter{},
		&GoogleDriveExporter{},
	}

	sut := NewRegisterExporter()

	// When

	exporters := sut.Register()

	// Then

	assert.Equal(t, expectedExporters, exporters)
}
