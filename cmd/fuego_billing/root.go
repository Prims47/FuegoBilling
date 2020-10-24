package main

import (
	"io"

	"github.com/prims47/FuegoBilling/internal/repository"
	"github.com/prims47/FuegoBilling/internal/services"
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
		NewGeneratePDFCmd(out, accountRepository, customerRepository, serviceRepository, formatFloat, formatInt),
	)

	return rootCmd, nil
}
