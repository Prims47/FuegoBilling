package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	generatedMock "github.com/prims47/FuegoBilling/cmd/fuego_billing/mock"
	"github.com/prims47/FuegoBilling/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePDFCmd(t *testing.T) {
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
			testName:                            "Given pass without account-config-path",
			expectedOutput:                      "../../tests/outputs/error_account_config_path_generate_pdf_command.txt",
			accountRepositoryMockRequestParams:  "",
			accountRepositoryMockTimes:          0,
			accountRepositoryMockModel:          model.Account{},
			accountRepositoryMockError:          nil,
			customerRepositoryMockRequestParams: "",
			customerRepositoryMockTimes:         0,
			customerRepositoryMockModel:         model.Customer{},
			customerRepositoryMockError:         nil,
			serviceRepositoryMockRequestParams:  "",
			serviceRepositoryMockTimes:          0,
			serviceRepositoryMockModel:          model.Service{},
			serviceRepositoryMockError:          nil,
			formatFloatMockRequestParams:        0,
			formatFloatMockTimes:                0,
			formatFloatMockReturn:               "",
			formatIntMockRequestParams1:         0,
			formatIntMockRequestParams2:         0,
			formatIntMockRequestParams3:         0,
			formatIntMockRequestParams4:         0,
			formatIntMockTimes:                  0,
			formatIntMockReturn:                 "",
		},
		{
			testName:                            "Given pass account-config-path without real value",
			expectedOutput:                      "../../tests/outputs/error_account_flag_config_path_generate_pdf_command.txt",
			accountRepositoryMockRequestParams:  "",
			accountRepositoryMockTimes:          0,
			accountRepositoryMockModel:          model.Account{},
			accountRepositoryMockError:          nil,
			customerRepositoryMockRequestParams: "",
			customerRepositoryMockTimes:         0,
			customerRepositoryMockModel:         model.Customer{},
			customerRepositoryMockError:         nil,
			serviceRepositoryMockRequestParams:  "",
			serviceRepositoryMockTimes:          0,
			serviceRepositoryMockModel:          model.Service{},
			serviceRepositoryMockError:          nil,
			formatFloatMockRequestParams:        0,
			formatFloatMockTimes:                0,
			formatFloatMockReturn:               "",
			formatIntMockRequestParams1:         0,
			formatIntMockRequestParams2:         0,
			formatIntMockRequestParams3:         0,
			formatIntMockRequestParams4:         0,
			formatIntMockTimes:                  0,
			formatIntMockReturn:                 "",
			args:                                []string{"--account-config-path"},
		},
		{
			testName:                            "Given pass without customer-config-path",
			expectedOutput:                      "../../tests/outputs/error_customer_config_path_generate_pdf_command.txt",
			accountRepositoryMockRequestParams:  "",
			accountRepositoryMockTimes:          0,
			accountRepositoryMockModel:          model.Account{},
			accountRepositoryMockError:          nil,
			customerRepositoryMockRequestParams: "",
			customerRepositoryMockTimes:         0,
			customerRepositoryMockModel:         model.Customer{},
			customerRepositoryMockError:         nil,
			serviceRepositoryMockRequestParams:  "",
			serviceRepositoryMockTimes:          0,
			serviceRepositoryMockModel:          model.Service{},
			serviceRepositoryMockError:          nil,
			formatFloatMockRequestParams:        0,
			formatFloatMockTimes:                0,
			formatFloatMockReturn:               "",
			formatIntMockRequestParams1:         0,
			formatIntMockRequestParams2:         0,
			formatIntMockRequestParams3:         0,
			formatIntMockRequestParams4:         0,
			formatIntMockTimes:                  0,
			formatIntMockReturn:                 "",
			args:                                []string{"--account-config-path", "../../tests/inputs/account.json"},
		},
		{
			testName:                            "Given pass customer-config-path without real value",
			expectedOutput:                      "../../tests/outputs/error_customer_flag_config_path_generate_pdf_command.txt",
			accountRepositoryMockRequestParams:  "",
			accountRepositoryMockTimes:          0,
			accountRepositoryMockModel:          model.Account{},
			accountRepositoryMockError:          nil,
			customerRepositoryMockRequestParams: "",
			customerRepositoryMockTimes:         0,
			customerRepositoryMockModel:         model.Customer{},
			customerRepositoryMockError:         nil,
			serviceRepositoryMockRequestParams:  "",
			serviceRepositoryMockTimes:          0,
			serviceRepositoryMockModel:          model.Service{},
			serviceRepositoryMockError:          nil,
			formatFloatMockRequestParams:        0,
			formatFloatMockTimes:                0,
			formatFloatMockReturn:               "",
			formatIntMockRequestParams1:         0,
			formatIntMockRequestParams2:         0,
			formatIntMockRequestParams3:         0,
			formatIntMockRequestParams4:         0,
			formatIntMockTimes:                  0,
			formatIntMockReturn:                 "",
			args:                                []string{"--account-config-path", "../../tests/inputs/account.json", "--customer-config-path"},
		},
		{
			testName:                            "Given pass without service-config-path",
			expectedOutput:                      "../../tests/outputs/error_service_config_path_generate_pdf_command.txt",
			accountRepositoryMockRequestParams:  "",
			accountRepositoryMockTimes:          0,
			accountRepositoryMockModel:          model.Account{},
			accountRepositoryMockError:          nil,
			customerRepositoryMockRequestParams: "",
			customerRepositoryMockTimes:         0,
			customerRepositoryMockModel:         model.Customer{},
			customerRepositoryMockError:         nil,
			serviceRepositoryMockRequestParams:  "",
			serviceRepositoryMockTimes:          0,
			serviceRepositoryMockModel:          model.Service{},
			serviceRepositoryMockError:          nil,
			formatFloatMockRequestParams:        0,
			formatFloatMockTimes:                0,
			formatFloatMockReturn:               "",
			formatIntMockRequestParams1:         0,
			formatIntMockRequestParams2:         0,
			formatIntMockRequestParams3:         0,
			formatIntMockRequestParams4:         0,
			formatIntMockTimes:                  0,
			formatIntMockReturn:                 "",
			args: []string{
				"--account-config-path",
				"../../tests/inputs/account.json",
				"--customer-config-path",
				"../../tests/inputs/customer.json",
			},
		},
		{
			testName:                            "Given pass service-config-path without real value",
			expectedOutput:                      "../../tests/outputs/error_service_flag_config_path_generate_pdf_command.txt",
			accountRepositoryMockRequestParams:  "",
			accountRepositoryMockTimes:          0,
			accountRepositoryMockModel:          model.Account{},
			accountRepositoryMockError:          nil,
			customerRepositoryMockRequestParams: "",
			customerRepositoryMockTimes:         0,
			customerRepositoryMockModel:         model.Customer{},
			customerRepositoryMockError:         nil,
			serviceRepositoryMockRequestParams:  "",
			serviceRepositoryMockTimes:          0,
			serviceRepositoryMockModel:          model.Service{},
			serviceRepositoryMockError:          nil,
			formatFloatMockRequestParams:        0,
			formatFloatMockTimes:                0,
			formatFloatMockReturn:               "",
			formatIntMockRequestParams1:         0,
			formatIntMockRequestParams2:         0,
			formatIntMockRequestParams3:         0,
			formatIntMockRequestParams4:         0,
			formatIntMockTimes:                  0,
			formatIntMockReturn:                 "",
			args: []string{
				"--account-config-path",
				"../../tests/inputs/account.json",
				"--customer-config-path",
				"../../tests/inputs/customer.json",
				"--service-config-path",
			},
		},
		{
			testName:                            "Given with AccountRepository with error",
			expectedOutput:                      "../../tests/outputs/error_account_config_path_generate_pdf_command.txt",
			accountRepositoryMockRequestParams:  "../../tests/inputs/account.json",
			accountRepositoryMockTimes:          1,
			accountRepositoryMockModel:          model.Account{},
			accountRepositoryMockError:          errors.New("It's not good my friends"),
			customerRepositoryMockRequestParams: "",
			customerRepositoryMockTimes:         0,
			customerRepositoryMockModel:         model.Customer{},
			customerRepositoryMockError:         nil,
			serviceRepositoryMockRequestParams:  "",
			serviceRepositoryMockTimes:          0,
			serviceRepositoryMockModel:          model.Service{},
			serviceRepositoryMockError:          nil,
			formatFloatMockRequestParams:        0,
			formatFloatMockTimes:                0,
			formatFloatMockReturn:               "",
			formatIntMockRequestParams1:         0,
			formatIntMockRequestParams2:         0,
			formatIntMockRequestParams3:         0,
			formatIntMockRequestParams4:         0,
			formatIntMockTimes:                  0,
			formatIntMockReturn:                 "",
			args: []string{
				"--account-config-path",
				"../../tests/inputs/account.json",
				"--customer-config-path",
				"../../tests/inputs/customer.json",
				"--service-config-path",
				"../../tests/inputs/service.json",
			},
		},
		{
			testName:                           "Given with CustomerRepository with error",
			expectedOutput:                     "../../tests/outputs/error_customer_config_path_generate_pdf_command.txt",
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
			customerRepositoryMockModel:         model.Customer{},
			customerRepositoryMockError:         errors.New("It's not good my friends"),
			serviceRepositoryMockRequestParams:  "",
			serviceRepositoryMockTimes:          0,
			serviceRepositoryMockModel:          model.Service{},
			serviceRepositoryMockError:          nil,
			formatFloatMockRequestParams:        0,
			formatFloatMockTimes:                0,
			formatFloatMockReturn:               "",
			formatIntMockRequestParams1:         0,
			formatIntMockRequestParams2:         0,
			formatIntMockRequestParams3:         0,
			formatIntMockRequestParams4:         0,
			formatIntMockTimes:                  0,
			formatIntMockReturn:                 "",
			args: []string{
				"--account-config-path",
				"../../tests/inputs/account.json",
				"--customer-config-path",
				"../../tests/inputs/customer.json",
				"--service-config-path",
				"../../tests/inputs/service.json",
			},
		},
		{
			testName:                           "Given with ServiceRepository with error",
			expectedOutput:                     "../../tests/outputs/error_service_config_path_generate_pdf_command.txt",
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
			serviceRepositoryMockModel:         model.Service{},
			serviceRepositoryMockError:         errors.New("It's not good my friends"),
			formatFloatMockRequestParams:       0,
			formatFloatMockTimes:               0,
			formatFloatMockReturn:              "",
			formatIntMockRequestParams1:        0,
			formatIntMockRequestParams2:        0,
			formatIntMockRequestParams3:        0,
			formatIntMockRequestParams4:        0,
			formatIntMockTimes:                 0,
			formatIntMockReturn:                "",
			args: []string{
				"--account-config-path",
				"../../tests/inputs/account.json",
				"--customer-config-path",
				"../../tests/inputs/customer.json",
				"--service-config-path",
				"../../tests/inputs/service.json",
			},
		},
		{
			testName:                           "Given generated PDF",
			expectedOutput:                     "",
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
				"../../tests/generated_pdf",
			},
		},
	}

	// When / Then

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

			buf := new(bytes.Buffer)

			sut := NewGeneratePDFCmd(buf, accountRepositoryMock, customerRepositoryMock, serviceRepositoryMock, formatFloatMock, formatIntMock)

			sut.SetOut(buf)
			sut.SetErr(buf)

			if len(tc.args) > 0 {
				sut.SetArgs(tc.args)
			}

			sut.ExecuteC()

			result := buf.String()

			if tc.expectedOutput != "" {
				expectedOutput, _ := ioutil.ReadFile(tc.expectedOutput)

				assert.Equal(t, []byte(result), expectedOutput)
			}

			if tc.expectedOutput == "" {
				if _, err := os.Stat("../../tests/generated_pdf"); err != nil {
					t.Fatal()
				}
			}
		})
	}

}

func deleteGeneratedPDF() {
	if _, err := os.Stat("../../tests/generated_pdf"); err == nil {
		os.RemoveAll("../../tests/generated_pdf")
	}
}
