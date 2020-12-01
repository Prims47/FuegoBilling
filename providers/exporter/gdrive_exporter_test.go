package providersExporter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSaveGoogleDriveExporter(t *testing.T) {
	t.Parallel()

	// Given

	testCases := []struct {
		testName       string
		canSaveParams  string
		expectedResult bool
	}{
		{testName: "Given empty provider name", canSaveParams: "", expectedResult: false},
		{testName: "Given unknow provider name", canSaveParams: "pepito", expectedResult: false},
		{testName: "Given lowercase google provider name", canSaveParams: "google", expectedResult: true},
		{testName: "Given first uppercase Google provider name", canSaveParams: "Google", expectedResult: true},
		{testName: "Given lowercase drive provider name", canSaveParams: "drive", expectedResult: true},
		{testName: "Given first uppercase Drive provider name", canSaveParams: "Drive", expectedResult: true},
	}

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			sut := &GoogleDriveExporter{}

			canSave := sut.CanSave(tc.canSaveParams)

			assert.Equal(t, tc.expectedResult, canSave)
		})
	}
}

func TestSaveGoogleDriveExporter(t *testing.T) {
	t.Parallel()

	// Given

	testCases := []struct {
		testName      string
		clientID      string
		clientSecret  string
		redirectURL   string
		authURL       string
		tokenURL      string
		expectedError string
	}{
		{
			testName:      "Given empty clientID",
			expectedError: "We must provide valid env var like 'DRIVE_CLIENT_ID', 'DRIVE_CLIENT_SECRET', 'DRIVE_REDIRECT_URL, DRIVE_AUTH_URL and DRIVE_TOKEN_URL'",
		},
		{
			testName:      "Given empty client secret",
			clientID:      "PepitoClientID",
			expectedError: "We must provide valid env var like 'DRIVE_CLIENT_ID', 'DRIVE_CLIENT_SECRET', 'DRIVE_REDIRECT_URL, DRIVE_AUTH_URL and DRIVE_TOKEN_URL'",
		},
		{
			testName:      "Given empty redirect url",
			clientID:      "PepitoClientID",
			clientSecret:  "PepitoClientSecret",
			expectedError: "We must provide valid env var like 'DRIVE_CLIENT_ID', 'DRIVE_CLIENT_SECRET', 'DRIVE_REDIRECT_URL, DRIVE_AUTH_URL and DRIVE_TOKEN_URL'",
		},
		{
			testName:      "Given empty auth url",
			clientID:      "PepitoClientID",
			clientSecret:  "PepitoClientSecret",
			redirectURL:   "PepitoRedirectURL",
			expectedError: "We must provide valid env var like 'DRIVE_CLIENT_ID', 'DRIVE_CLIENT_SECRET', 'DRIVE_REDIRECT_URL, DRIVE_AUTH_URL and DRIVE_TOKEN_URL'",
		},
		{
			testName:      "Given empty token url",
			clientID:      "PepitoClientID",
			clientSecret:  "PepitoClientSecret",
			redirectURL:   "PepitoRedirectURL",
			authURL:       "PepitoAuthURL",
			expectedError: "We must provide valid env var like 'DRIVE_CLIENT_ID', 'DRIVE_CLIENT_SECRET', 'DRIVE_REDIRECT_URL, DRIVE_AUTH_URL and DRIVE_TOKEN_URL'",
		},
	}

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {

			os.Setenv("DRIVE_CLIENT_ID", tc.clientID)
			os.Setenv("DRIVE_CLIENT_SECRET", tc.clientSecret)
			os.Setenv("DRIVE_REDIRECT_URL", tc.redirectURL)
			os.Setenv("DRIVE_AUTH_URL", tc.authURL)
			os.Setenv("DRIVE_TOKEN_URL", tc.tokenURL)

			sut := &GoogleDriveExporter{}

			err := sut.Save("pepito.txt", []byte("hello world"))

			assert.Equal(t, tc.expectedError, err.Error())
		})
	}

}
