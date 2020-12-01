package providersExporter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanSaveAWSExporter(t *testing.T) {
	t.Parallel()

	// Given

	testCases := []struct {
		testName       string
		canSaveParams  string
		expectedResult bool
	}{
		{testName: "Given empty provider name", canSaveParams: "", expectedResult: false},
		{testName: "Given unknow provider name", canSaveParams: "pepito", expectedResult: false},
		{testName: "Given lowercase aws provider name", canSaveParams: "aws", expectedResult: true},
		{testName: "Given uppercase AWS provider name", canSaveParams: "AWS", expectedResult: true},
	}

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			sut := &AWSExporter{}

			canSave := sut.CanSave(tc.canSaveParams)

			assert.Equal(t, tc.expectedResult, canSave)

		})
	}
}

func TestSaveAWSExporter(t *testing.T) {
	t.Parallel()

	// Given

	testCases := []struct {
		testName      string
		accesKey      string
		secretKey     string
		region        string
		bucket        string
		expectedError string
	}{
		{
			testName:      "Given empty access key",
			secretKey:     "pepitoSecret",
			region:        "mexico",
			bucket:        "fuego",
			expectedError: "We must provide valid env var like 'AWS_ACCESS_KEY_ID', 'AWS_SECRET_ACCESS_KEY', 'AWS_REGION and AWS_BUCKET_NAME'",
		},
		{
			testName:      "Given empty secret key",
			accesKey:      "pepitoAccess",
			region:        "mexico",
			bucket:        "fuego",
			expectedError: "We must provide valid env var like 'AWS_ACCESS_KEY_ID', 'AWS_SECRET_ACCESS_KEY', 'AWS_REGION and AWS_BUCKET_NAME'",
		},
		{
			testName:      "Given empty region",
			accesKey:      "pepitoAccess",
			secretKey:     "pepitoSecret",
			bucket:        "fuego",
			expectedError: "We must provide valid env var like 'AWS_ACCESS_KEY_ID', 'AWS_SECRET_ACCESS_KEY', 'AWS_REGION and AWS_BUCKET_NAME'",
		},
		{
			testName:      "Given empty bucket",
			accesKey:      "pepitoAccess",
			secretKey:     "pepitoSecret",
			region:        "mexico",
			expectedError: "We must provide valid env var like 'AWS_ACCESS_KEY_ID', 'AWS_SECRET_ACCESS_KEY', 'AWS_REGION and AWS_BUCKET_NAME'",
		},
		{
			testName:      "Given wrong creds",
			accesKey:      "pepitoAccess",
			secretKey:     "pepitoSecret",
			region:        "mexico",
			bucket:        "fuego",
			expectedError: "RequestError: send request failed\ncaused by: Put \"https://fuego.s3.mexico.amazonaws.com/pepito.txt\": dial tcp: lookup fuego.s3.mexico.amazonaws.com: no such host",
		},
	}

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {

			os.Setenv("AWS_ACCESS_KEY_ID", tc.accesKey)
			os.Setenv("AWS_SECRET_ACCESS_KEY", tc.secretKey)
			os.Setenv("AWS_REGION", tc.region)
			os.Setenv("AWS_BUCKET_NAME", tc.bucket)

			sut := &AWSExporter{}

			err := sut.Save("pepito.txt", []byte("hello world"))

			assert.Equal(t, tc.expectedError, err.Error())
		})
	}

}
