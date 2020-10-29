package providersExporter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSaveLocalExporter(t *testing.T) {
	t.Parallel()

	// Given

	testCases := []struct {
		testName       string
		canSaveParams  string
		expectedResult bool
	}{
		{testName: "Given empty provider name", canSaveParams: "", expectedResult: true},
		{testName: "Given unknow provider name", canSaveParams: "pepito", expectedResult: false},
		{testName: "Given locale provider name", canSaveParams: "local", expectedResult: true},
	}

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			sut := &LocalExporter{}

			canSave := sut.CanSave(tc.canSaveParams)

			assert.Equal(t, tc.expectedResult, canSave)

		})
	}
}

func TestSaveLocalExporter(t *testing.T) {
	t.Parallel()

	// Given

	testCases := []struct {
		testName       string
		fileName       string
		saveData       []byte
		fileSize       int64
		expectedResult error
	}{
		{testName: "Given save data", fileName: "pepito.pdf", saveData: []byte("Pepito Fuego"), fileSize: 12, expectedResult: nil},
	}

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			defer deleteFolder()

			sut := &LocalExporter{}

			err := sut.Save(tc.fileName, tc.saveData)

			assert.Equal(t, tc.expectedResult, err)

			if tc.expectedResult == nil {
				fileInfo, _ := os.Stat("generated-pdf/" + tc.fileName)

				assert.Equal(t, tc.fileSize, fileInfo.Size())
			}

		})
	}
}

func deleteFolder() {
	if _, err := os.Stat("generated-pdf"); err == nil {
		os.RemoveAll("generated-pdf")
	}
}
