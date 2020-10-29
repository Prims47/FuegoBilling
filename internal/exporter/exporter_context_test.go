package exporter

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	generatedMock "github.com/prims47/FuegoBilling/internal/exporter/mock"
	"github.com/stretchr/testify/assert"
)

func TestExpoterContext(t *testing.T) {
	t.Parallel()

	// Given

	mockCtrl := gomock.NewController(t)

	exporterProviderMock := generatedMock.NewMockExporterProviderInterface(mockCtrl)

	testCases := []struct {
		testName              string
		fileName              string
		exporterCanSaveParams string
		exporterCanSaveTimes  int
		exporterCanSaveReturn bool
		exporterSaveParams    []byte
		exporterSaveTimes     int
		exporterSaveReturn    error
		saveParams            string
		saveReturn            error
	}{
		{
			testName:              "Given save without params",
			fileName:              "pepito.pdf",
			saveParams:            "",
			saveReturn:            errors.New("No provider found"),
			exporterCanSaveParams: "",
			exporterCanSaveTimes:  1,
			exporterCanSaveReturn: false,
			exporterSaveParams:    []byte(""),
			exporterSaveTimes:     0,
			exporterSaveReturn:    nil,
		},
		{
			testName:              "Given save without good params",
			fileName:              "pepito.pdf",
			saveParams:            "fuegobyp",
			saveReturn:            errors.New("No provider found"),
			exporterCanSaveParams: "fuegobyp",
			exporterCanSaveTimes:  1,
			exporterCanSaveReturn: false,
			exporterSaveParams:    []byte(""),
			exporterSaveTimes:     0,
			exporterSaveReturn:    nil,
		},
		{
			testName:              "Given save with good params",
			saveParams:            "aws",
			saveReturn:            nil,
			exporterCanSaveParams: "aws",
			exporterCanSaveTimes:  1,
			exporterCanSaveReturn: true,
			exporterSaveParams:    []byte("hello"),
			exporterSaveTimes:     1,
			exporterSaveReturn:    nil,
		},
		{
			testName:              "Given save with good params and error when provider save",
			saveParams:            "aws",
			saveReturn:            errors.New("No esta bueno my amigo!"),
			exporterCanSaveParams: "aws",
			exporterCanSaveTimes:  1,
			exporterCanSaveReturn: true,
			exporterSaveParams:    []byte("hello"),
			exporterSaveTimes:     1,
			exporterSaveReturn:    errors.New("No esta bueno my amigo!"),
		},
	}

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			exporterProviderMock.EXPECT().CanSave(tc.exporterCanSaveParams).Times(tc.exporterCanSaveTimes).Return(tc.exporterCanSaveReturn)
			exporterProviderMock.EXPECT().Save(tc.fileName, tc.exporterSaveParams).Times(tc.exporterSaveTimes).Return(tc.exporterSaveReturn)

			providers := []ExporterProviderInterface{exporterProviderMock}

			sut := NewExporterProviderContext(providers)

			err := sut.Save(tc.fileName, tc.saveParams, []byte("hello"))

			assert.Equal(t, tc.saveReturn, err)
		})
	}
}
