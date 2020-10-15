package main

import (
	"fmt"
	"io"

	"fuegobyp-billing.com/internal/repository"
	"fuegobyp-billing.com/internal/services"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const desc = `ffkfkfkf`

func NewCreatePdfCmd(out io.Writer,
	accountRepository repository.AccountRepositoryInterface,
	customerRepository repository.CustomerRepositoryInterface,
	serviceRepository repository.ServiceRepositoryInterface,
	formatFloat services.FormatFloatInterface,
	formatInt services.FormatIntInterface) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-pdf",
		Short:   "Generate billing",
		Long:    desc,
		Aliases: []string{"cpdf"},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) == 0 {
				fmt.Println("ShellCompDirectiveDefault")
				return nil, cobra.ShellCompDirectiveDefault
			}

			fmt.Println("ShellCompDirectiveNoFileComp")
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			configAccountPath, err := cmd.Flags().GetString("account-config-path")

			if err != nil || configAccountPath == "" {
				return errors.Errorf("Please give a valid account config path")
			}

			configCustomerPath, err := cmd.Flags().GetString("customer-config-path")

			if err != nil || configCustomerPath == "" {
				return errors.Errorf("Please give a valid customer config path")
			}

			configServicePath, err := cmd.Flags().GetString("service-config-path")

			if err != nil || configServicePath == "" {
				return errors.Errorf("Please give a valid service config path")
			}

			account, err := accountRepository.Request(configAccountPath)

			if err != nil {
				return errors.Errorf("Please give a valid account config path")
			}

			customer, err := customerRepository.Request(configCustomerPath)

			if err != nil {
				return errors.Errorf("Please give a valid customer config path")
			}

			service, err := serviceRepository.Request(configServicePath)

			if err != nil {
				return errors.Errorf("Please give a valid service config path")
			}

			// billing := pdf.NewBillingPDF("pdf", "toto", account, customer, service, formatInt, formatFloat)
			// billing.CreatePdf()

			fmt.Println(account)
			fmt.Println(customer)
			fmt.Println(service)

			return nil
		},
	}

	cmd.Flags().StringP("account-config-path", "a", "", "JSON Account Config Path")
	cmd.Flags().StringP("customer-config-path", "c", "", "JSON Customer Config Path")
	cmd.Flags().StringP("service-config-path", "s", "", "JSON Service Config Path")

	return cmd
}
