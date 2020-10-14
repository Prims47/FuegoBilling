package main

import (
	"io"

	"fuegobyp-billing.com/internal/repository"
	"github.com/spf13/cobra"
)

func newRootCmd(out io.Writer,
	accountRepository repository.AccountRepositoryInterface,
	customerRepository repository.CustomerRepositoryInterface) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "fuegoBilling",
		Short: "fuegoBilling is a very fast tool billing generator",
	}

	rootCmd.AddCommand(
		NewCreatePdfCmd(out, accountRepository, customerRepository),
	)

	return rootCmd, nil
}
