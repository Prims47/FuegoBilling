package pdf

import (
	"os"
	"testing"
	"time"

	"fuegobyp-billing.com/internal/model"
	generatedMock "fuegobyp-billing.com/internal/pdf/mock"
	"github.com/golang/mock/gomock"
	"github.com/jung-kurt/gofpdf"
)

const generatedPDFPath = "../../tests/generated_pdf"

func TestGeneratePDF(t *testing.T) {
	t.Parallel()

	// Given

	mockCtrl := gomock.NewController(t)

	accountRepositoryMock := generatedMock.NewMockAccountRepositoryInterface(mockCtrl)
	customerRepositoryMock := generatedMock.NewMockCustomerRepositoryInterface(mockCtrl)
	serviceRepositoryMock := generatedMock.NewMockServiceRepositoryInterface(mockCtrl)
	formatFloatMock := generatedMock.NewMockFormatFloatInterface(mockCtrl)
	formatIntMock := generatedMock.NewMockFormatIntInterface(mockCtrl)

	testCases := []struct {
		testName                            string
		pdfName                             string
		expectedOutput                      string
		accountRepositoryMockRequestParams  string
		accountRepositoryMockTimes          int
		accountRepositoryMockModel          model.Account
		accountRepositoryMockError          error
		customerRepositoryMockRequestParams string
		customerRepositoryMockTimes         int
		customerRepositoryMockModel         model.Customer
		customerRepositoryMockError         error
		serviceRepositoryMockRequestParams  string
		serviceRepositoryMockTimes          int
		serviceRepositoryMockModel          model.Service
		serviceRepositoryMockError          error
		formatFloatMockRequestParams        float32
		formatFloatMockTimes                int
		formatFloatMockReturn               string
		formatIntMockRequestParams1         int
		formatIntMockRequestParams2         int
		formatIntMockRequestParams3         int
		formatIntMockRequestParams4         int
		formatIntMockTimes                  int
		formatIntMockReturn                 string
		args                                []string
	}{
		{
			testName:                           "Given generated PDF",
			pdfName:                            "test_generated_billing",
			expectedOutput:                     "../../tests/pdf/test_generated_billing.pdf",
			accountRepositoryMockRequestParams: "../../tests/inputs/account.json",
			accountRepositoryMockTimes:         1,
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
			accountRepositoryMockError:          nil,
			customerRepositoryMockRequestParams: "../../tests/inputs/customer.json",
			customerRepositoryMockTimes:         1,
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
			customerRepositoryMockError:        nil,
			serviceRepositoryMockRequestParams: "../../tests/inputs/service.json",
			serviceRepositoryMockTimes:         1,
			serviceRepositoryMockModel: model.Service{
				Detail:    "Prestation Pepito Fuego by P",
				Quantity:  10,
				UnitPrice: 663,
				TVA:       model.TVA{Pourcent: 20},
			},
			serviceRepositoryMockError:   nil,
			formatFloatMockRequestParams: 20,
			formatFloatMockTimes:         1,
			formatFloatMockReturn:        "20",
			formatIntMockRequestParams1:  6630,
			formatIntMockRequestParams2:  6630,
			formatIntMockRequestParams3:  1326,
			formatIntMockRequestParams4:  7956,
			formatIntMockTimes:           1,
			formatIntMockReturn:          "38523",
			args: []string{
				"--account-config-path",
				"../../tests/inputs/account.json",
				"--customer-config-path",
				"../../tests/inputs/customer.json",
				"--service-config-path",
				"../../tests/inputs/service.json",
				"--pdf-path",
				generatedPDFPath,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			defer deleteGeneratedPDF()

			accountRepositoryMock.EXPECT().
				Request(gomock.Eq(tc.accountRepositoryMockRequestParams)).
				Times(tc.accountRepositoryMockTimes).
				Return(tc.accountRepositoryMockModel, tc.accountRepositoryMockError)

			customerRepositoryMock.EXPECT().
				Request(gomock.Eq(tc.customerRepositoryMockRequestParams)).
				Times(tc.customerRepositoryMockTimes).
				Return(tc.customerRepositoryMockModel, tc.customerRepositoryMockError)

			serviceRepositoryMock.EXPECT().
				Request(gomock.Eq(tc.serviceRepositoryMockRequestParams)).
				Times(tc.serviceRepositoryMockTimes).
				Return(tc.serviceRepositoryMockModel, tc.serviceRepositoryMockError)

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

			sut := NewBillingPDF(
				generatedPDFPath,
				tc.pdfName,
				tc.accountRepositoryMockModel,
				tc.customerRepositoryMockModel,
				tc.serviceRepositoryMockModel,
				formatIntMock,
				formatFloatMock,
				"FR-1234",
				"22 Oct 2020",
			)

			os.Mkdir(generatedPDFPath, os.ModePerm)

			datePDF := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
			gofpdf.SetDefaultCreationDate(datePDF)
			gofpdf.SetDefaultModificationDate(datePDF)

			sut.CreatePDF()

			if tc.expectedOutput != "" {
				os.Chtimes(tc.expectedOutput, datePDF, datePDF)
				err := gofpdf.ComparePDFFiles(tc.expectedOutput, generatedPDFPath+"/test_generated_billing.pdf", true)

				if err != nil {
					t.Fatal("Error when compare expected PDF")
				}
			}

			if tc.expectedOutput == "" {
				if _, err := os.Stat(generatedPDFPath); err != nil {
					t.Fatal()
				}
			}
		})
	}

}

func deleteGeneratedPDF() {
	if _, err := os.Stat(generatedPDFPath); err == nil {
		os.RemoveAll(generatedPDFPath)
	}
}
