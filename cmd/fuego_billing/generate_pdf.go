package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"fuegobyp-billing.com/internal/model"
	"fuegobyp-billing.com/internal/pdf"
	"fuegobyp-billing.com/internal/repository"
	"fuegobyp-billing.com/internal/services"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const desc = `Generate billing easly`
const dateFormat = "02-01-2006"

func NewGeneratePDFCmd(out io.Writer,
	accountRepository repository.AccountRepositoryInterface,
	customerRepository repository.CustomerRepositoryInterface,
	serviceRepository repository.ServiceRepositoryInterface,
	formatFloat services.FormatFloatInterface,
	formatInt services.FormatIntInterface) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate-pdf",
		Short:   "Generate billing",
		Long:    desc,
		Aliases: []string{"gpdf"},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				fmt.Println("ShellCompDirectiveDefault")
				return nil, cobra.ShellCompDirectiveDefault
			}

			fmt.Println("ShellCompDirectiveNoFileComp")
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
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
			pdfPath := handlePDFPath(cmd)

			billingPDF := pdf.NewBillingPDF(pdfPath, pdfName, account, customer, service, formatInt, formatFloat, billingNumber)
			billingPDF.CreatePdf()

			return nil
		},
	}

	cmd.Flags().StringP("account-config-path", "a", "", "JSON Account Config Path")
	cmd.Flags().StringP("customer-config-path", "c", "", "JSON Customer Config Path")
	cmd.Flags().StringP("service-config-path", "s", "", "JSON Service Config Path")
	cmd.Flags().StringP("pdf-path", "p", "", "PDF Path")

	return cmd
}

func cleanPDFName(billingNumber string, customerName string) string {
	return fmt.Sprintf("billing-%s-customer-%s-date-to-%s", billingNumber, strings.Replace(strings.ToLower(customerName), " ", "-", -1), time.Now().Format(dateFormat))
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

func handlePDFPath(cmd *cobra.Command) string {
	pdfPath, err := cmd.Flags().GetString("pdf-path")

	if err != nil || pdfPath == "" {
		pdfPath = "pdf"
	}

	if _, err := os.Stat(pdfPath); err != nil {
		os.Mkdir(pdfPath, os.ModePerm)
	}

	return pdfPath
}
