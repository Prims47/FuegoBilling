package pdf

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jung-kurt/gofpdf"
	"github.com/prims47/FuegoBilling/internal/model"
	generatedMock "github.com/prims47/FuegoBilling/internal/pdf/mock"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePDF(t *testing.T) {
	t.Parallel()

	// Given

	mockCtrl := gomock.NewController(t)

	formatFloatMock := generatedMock.NewMockFormatFloatInterface(mockCtrl)
	formatIntMock := generatedMock.NewMockFormatIntInterface(mockCtrl)

	testCases := []struct {
		testName                     string
		pdfName                      string
		expectedOutput               string
		accountRepositoryMockModel   model.Account
		customerRepositoryMockModel  model.Customer
		serviceRepositoryMockModel   model.Service
		formatFloatMockRequestParams float32
		formatFloatMockTimes         int
		formatFloatMockReturn        string
		formatIntMockRequestParams1  int
		formatIntMockRequestParams2  int
		formatIntMockRequestParams3  int
		formatIntMockRequestParams4  int
		formatIntMockTimes           int
		formatIntMockReturn          string
	}{
		{
			testName:       "Given generated PDF",
			expectedOutput: "../../tests/pdf/test_generated_billing.pdf",
			accountRepositoryMockModel: model.Account{
				Name:      "Pepito",
				FirstName: "Ilan",
				LastName:  "Zerath",
				Mail:      "pepito@fuegobyp.io",
				Address: model.Address{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "France",
				},
				Company: model.Company{
					Siret:   "11212",
					Tva:     "21212",
					Capital: 100.47,
					RCS:     "Paris",
					NAF:     "NAF",
					Type:    "SARL",
				},
			},
			customerRepositoryMockModel: model.Customer{
				Name: "Tesla",
				Address: model.Address{
					Street:  "770 rue du Fuego",
					ZipCode: "75006",
					City:    "Paris",
					Country: "France",
				},
				Company: model.Company{
					Siret:   "11212",
					Tva:     "21212",
					Capital: 0,
					RCS:     "",
					NAF:     "",
					Type:    "SARL",
				},
			},
			serviceRepositoryMockModel: model.Service{
				Detail:    "Prestation Pepito Fuego by P",
				Quantity:  10,
				UnitPrice: 663,
				TVA:       model.TVA{Pourcent: 20},
			},
			formatFloatMockRequestParams: 20,
			formatFloatMockTimes:         1,
			formatFloatMockReturn:        "20",
			formatIntMockRequestParams1:  6630,
			formatIntMockRequestParams2:  6630,
			formatIntMockRequestParams3:  1326,
			formatIntMockRequestParams4:  7956,
			formatIntMockTimes:           1,
			formatIntMockReturn:          "38523",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {

			formatFloatMock.EXPECT().
				Float32ToString(tc.formatFloatMockRequestParams).
				Times(tc.formatFloatMockTimes).
				Return(tc.formatFloatMockReturn)

			gomock.InOrder(
				formatIntMock.EXPECT().
					IntToStringFrenchFormat(tc.formatIntMockRequestParams1).
					Times(tc.formatIntMockTimes).
					Return(tc.formatIntMockReturn),
				formatIntMock.EXPECT().
					IntToStringFrenchFormat(tc.formatIntMockRequestParams2).
					Times(tc.formatIntMockTimes).
					Return(tc.formatIntMockReturn),
				formatIntMock.EXPECT().
					IntToStringFrenchFormat(tc.formatIntMockRequestParams3).
					Times(tc.formatIntMockTimes).
					Return(tc.formatIntMockReturn),
				formatIntMock.EXPECT().
					IntToStringFrenchFormat(tc.formatIntMockRequestParams4).
					Times(tc.formatIntMockTimes).
					Return(tc.formatIntMockReturn),
			)

			buf := new(bytes.Buffer)

			sut := NewBillingPDF(
				tc.accountRepositoryMockModel,
				tc.customerRepositoryMockModel,
				tc.serviceRepositoryMockModel,
				formatIntMock,
				formatFloatMock,
				"FR-1234",
				"22 Oct 2020",
				buf,
			)

			datePDF := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			gofpdf.SetDefaultCreationDate(datePDF)
			gofpdf.SetDefaultModificationDate(datePDF)

			sut.CreatePDF()

			if tc.expectedOutput != "" {
				os.Chtimes(tc.expectedOutput, datePDF, datePDF)

				data, err := ioutil.ReadFile(tc.expectedOutput)

				if err != nil {
					t.Fatal()
				}

				assert.Equal(t, data, buf.Bytes())
			}
		})
	}

}
