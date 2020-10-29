package repository

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	generatedMock "github.com/prims47/FuegoBilling/internal/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestExporterRepository(t *testing.T) {
	t.Parallel()

	// Given

	mockCtrl := gomock.NewController(t)

	exporterContextMock := generatedMock.NewMockExporterContextInterface(mockCtrl)

	testCases := []struct {
		testName           string
		fileName           string
		mockSaveParamsName string
		mockSaveParamsData []byte
		mockSaveTimes      int
		mockSaveReturn     error
		saveParamsName     string
		saveParamsData     []byte
		expectSaveReturn   error
	}{
		{
			testName:           "Given empty exporter name",
			fileName:           "pepito.fuego.pdf",
			mockSaveParamsName: "",
			mockSaveParamsData: []byte(""),
			mockSaveTimes:      1,
			mockSaveReturn:     errors.New("No provider found"),
			saveParamsName:     "",
			saveParamsData:     []byte(""),
			expectSaveReturn:   errors.New("No provider found"),
		},
		{
			testName:           "Given incorrect exporter name",
			fileName:           "pepito.fuego.pdf",
			mockSaveParamsName: "BlaBla",
			mockSaveParamsData: []byte("My super PDF"),
			mockSaveTimes:      1,
			mockSaveReturn:     errors.New("No provider found"),
			saveParamsName:     "BlaBla",
			saveParamsData:     []byte("My super PDF"),
			expectSaveReturn:   errors.New("No provider found"),
		},
		{
			testName:           "Given valid exporter",
			fileName:           "pepito.fuego.pdf",
			mockSaveParamsName: "AWS",
			mockSaveParamsData: []byte("My super PDF"),
			mockSaveTimes:      1,
			mockSaveReturn:     nil,
			saveParamsName:     "AWS",
			saveParamsData:     []byte("My super PDF"),
			expectSaveReturn:   nil,
		},
	}

	sut := NewExporterRepositoryRepository(exporterContextMock)

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			exporterContextMock.EXPECT().Save(tc.fileName, tc.mockSaveParamsName, tc.mockSaveParamsData).Times(tc.mockSaveTimes).Return(tc.mockSaveReturn)

			err := sut.Save(tc.fileName, tc.saveParamsName, tc.saveParamsData)

			assert.Equal(t, tc.expectSaveReturn, err)
		})
	}
}
