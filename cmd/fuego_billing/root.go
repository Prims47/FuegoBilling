package main

import (
	"io"

	"fuegobyp-billing.com/internal/repository"
	"fuegobyp-billing.com/internal/services"
	"github.com/spf13/cobra"
)

func newRootCmd(out io.Writer,
	accountRepository repository.AccountRepositoryInterface,
	customerRepository repository.CustomerRepositoryInterface,
	serviceRepository repository.ServiceRepositoryInterface,
	formatFloat services.FormatFloatInterface,
	formatInt services.FormatIntInterface) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "fuegoBilling",
		Short: "fuegoBilling is a very fast tool billing generator",
	}

	rootCmd.AddCommand(
		NewCreatePdfCmd(out, accountRepository, customerRepository, serviceRepository, formatFloat, formatInt),
	)

	return rootCmd, nil
}
