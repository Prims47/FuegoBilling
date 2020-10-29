package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/prims47/FuegoBilling/internal/exporter"
	"github.com/prims47/FuegoBilling/internal/model"
	"github.com/prims47/FuegoBilling/internal/pdf"
	"github.com/prims47/FuegoBilling/internal/repository"
	"github.com/prims47/FuegoBilling/internal/services"
	"github.com/spf13/cobra"
)

const desc = `Generate billing easly`
const dateFormat = "02-01-2006"
const dateFormatToPDF = "2 Jan, 2006"

func NewGeneratePDFCmd(out io.Writer,
	accountRepository repository.AccountRepositoryInterface,
	customerRepository repository.CustomerRepositoryInterface,
	serviceRepository repository.ServiceRepositoryInterface,
	formatFloat services.FormatFloatInterface,
	formatInt services.FormatIntInterface,
	exporterContext exporter.ExporterContextInterface) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate-pdf",
		Short:   "Generate billing",
		Long:    desc,
		Aliases: []string{"gpdf"},
		RunE: func(cmd *cobra.Command, args []string) error {
			configAccountPath, configCustomerPath, configServicePath, err := handleConfigs(cmd)

			if err != nil {
				return err
			}

			account, customer, service, err := handleRepositories(accountRepository,
				customerRepository,
				serviceRepository,
				configAccountPath,
				configCustomerPath,
				configServicePath,
			)

			if err != nil {
				return err
			}

			billing := model.Billing{}
			billingNumber := billing.GetBillingNumber()

			pdfName := cleanPDFName(billingNumber, customer.Name)

			buf := new(bytes.Buffer)

			billingPDF := pdf.NewBillingPDF(account, customer, service, formatInt, formatFloat, billingNumber, time.Now().Format(dateFormatToPDF), buf)
			billingPDF.CreatePDF()

			exportFlag, _ := cmd.Flags().GetString("export")

			exporterError := exporterContext.Save(pdfName, exportFlag, buf.Bytes())

			if exporterError != nil {
				return exporterError
			}

			return nil
		},
	}

	cmd.Flags().StringP("account-config-path", "a", "", "JSON Account Config Path")
	cmd.Flags().StringP("customer-config-path", "c", "", "JSON Customer Config Path")
	cmd.Flags().StringP("service-config-path", "s", "", "JSON Service Config Path")
	cmd.Flags().StringP("export", "e", "", "Exporter provider (ex: local, AWS)")

	return cmd
}

func cleanPDFName(billingNumber string, customerName string) string {
	return fmt.Sprintf("billing-%s-customer-%s-date-to-%s.pdf", billingNumber, strings.Replace(strings.ToLower(customerName), " ", "-", -1), time.Now().Format(dateFormat))
}

func handleConfigs(cmd *cobra.Command) (string, string, string, error) {
	configAccountPath, err := cmd.Flags().GetString("account-config-path")

	if err != nil || configAccountPath == "" {
		return "", "", "", errors.Errorf("Please give a valid account config path")
	}

	configCustomerPath, err := cmd.Flags().GetString("customer-config-path")

	if err != nil || configCustomerPath == "" {
		return "", "", "", errors.Errorf("Please give a valid customer config path")
	}

	configServicePath, err := cmd.Flags().GetString("service-config-path")

	if err != nil || configServicePath == "" {
		return "", "", "", errors.Errorf("Please give a valid service config path")
	}

	return configAccountPath, configCustomerPath, configServicePath, nil
}

func handleRepositories(accountRepository repository.AccountRepositoryInterface,
	customerRepository repository.CustomerRepositoryInterface,
	serviceRepository repository.ServiceRepositoryInterface,
	configAccountPath string,
	configCustomerPath string,
	configServicePath string) (model.Account, model.Customer, model.Service, error) {
	account, err := accountRepository.Request(configAccountPath)

	if err != nil {
		return model.Account{}, model.Customer{}, model.Service{}, errors.Errorf("Please give a valid account config path")
	}

	customer, err := customerRepository.Request(configCustomerPath)

	if err != nil {
		return model.Account{}, model.Customer{}, model.Service{}, errors.Errorf("Please give a valid customer config path")
	}

	service, err := serviceRepository.Request(configServicePath)

	if err != nil {
		return model.Account{}, model.Customer{}, model.Service{}, errors.Errorf("Please give a valid service config path")
	}

	return account, customer, service, nil
}
