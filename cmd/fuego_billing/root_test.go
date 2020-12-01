package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	generatedMock "github.com/prims47/FuegoBilling/cmd/fuego_billing/mock"
	"github.com/prims47/FuegoBilling/internal/model"
	"github.com/prims47/FuegoBilling/internal/repository"
	"github.com/prims47/FuegoBilling/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestRootCmd(t *testing.T) {
	t.Parallel()

	// Given

	mockCtrl := gomock.NewController(t)

	accountRepositoryMock := generatedMock.NewMockAccountRepositoryInterface(mockCtrl)
	customerRepositoryMock := generatedMock.NewMockCustomerRepositoryInterface(mockCtrl)
	serviceRepositoryMock := generatedMock.NewMockServiceRepositoryInterface(mockCtrl)
	formatFloatMock := generatedMock.NewMockFormatFloatInterface(mockCtrl)
	formatIntMock := generatedMock.NewMockFormatIntInterface(mockCtrl)
	exporterMock := generatedMock.NewMockExporterContextInterface(mockCtrl)

	accountRepositoryMock.EXPECT().
		Request(gomock.Eq("")).
		Times(0).
		Return(model.Account{}, nil)

	customerRepositoryMock.EXPECT().
		Request(gomock.Eq("")).
		Times(0).
		Return(model.Customer{}, nil)

	serviceRepositoryMock.EXPECT().
		Request(gomock.Eq("")).
		Times(0).
		Return(model.Service{}, nil)

	formatFloatMock.EXPECT().
		Float32ToString(gomock.Eq("")).
		Times(0).
		Return("")

	formatIntMock.EXPECT().
		IntToStringFrenchFormat(gomock.Eq("")).
		Times(0).
		Return("")

	exporterMock.EXPECT().
		Save(gomock.Eq(""), gomock.Eq(""), gomock.Eq([]byte(""))).
		Times(0).
		Return(nil)

	testCases := []struct {
		testName               string
		expectedOutput         string
		accountRepositoryMock  repository.AccountRepositoryInterface
		customerRepositoryMock repository.CustomerRepositoryInterface
		serviceRepositoryMock  repository.ServiceRepositoryInterface
		formatFloatMock        services.FormatFloatInterface
		formatIntMock          services.FormatIntInterface
		args                   []string
	}{
		{
			testName:               "Given pass no args",
			expectedOutput:         "../../tests/outputs/root_command.txt",
			accountRepositoryMock:  accountRepositoryMock,
			customerRepositoryMock: customerRepositoryMock,
			serviceRepositoryMock:  serviceRepositoryMock,
			formatFloatMock:        formatFloatMock,
			formatIntMock:          formatIntMock,
		},
		{
			testName:               "Given pass wrong args",
			expectedOutput:         "../../tests/outputs/root_command_error.txt",
			accountRepositoryMock:  accountRepositoryMock,
			customerRepositoryMock: customerRepositoryMock,
			serviceRepositoryMock:  serviceRepositoryMock,
			formatFloatMock:        formatFloatMock,
			formatIntMock:          formatIntMock,
			args:                   []string{"pepito"},
		},
	}

	// When / Then

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			buf := new(bytes.Buffer)

			sut, err := newRootCmd(buf, tc.accountRepositoryMock, tc.customerRepositoryMock, tc.serviceRepositoryMock, tc.formatFloatMock, tc.formatIntMock, exporterMock)

			sut.SetOut(buf)
			sut.SetErr(buf)

			if len(tc.args) > 0 {
				sut.SetArgs(tc.args)
			}

			sut.ExecuteC()

			result := buf.String()

			expectedOutput, _ := ioutil.ReadFile(tc.expectedOutput)

			assert.True(t, bytes.Equal([]byte(result), expectedOutput))
			assert.Empty(t, err)
		})
	}

}
