package main

import (
	"fmt"
	"os"

	"github.com/prims47/FuegoBilling/internal/adapter"
	"github.com/prims47/FuegoBilling/internal/exporter"
	"github.com/prims47/FuegoBilling/internal/repository"
	"github.com/prims47/FuegoBilling/internal/services"
	providersExporter "github.com/prims47/FuegoBilling/providers/exporter"
)

func main() {
	accountAdapter := adapter.NewAccountJSONAdapter()
	accountRepository := repository.NewAccountRepository(accountAdapter)

	customerAdapter := adapter.NewCustomerJSONAdapter()
	customerRepository := repository.NewCustomerRepository(customerAdapter)

	serviceAdapter := adapter.NewServiceJSONAdapter()
	serviceRepository := repository.NewServiceRepository(serviceAdapter)

	registerExporter := providersExporter.NewRegisterExporter()
	exporters := registerExporter.Register()

	exporterContext := exporter.NewExporterProviderContext(exporters)

	cmd, err := newRootCmd(
		os.Stdout,
		accountRepository,
		customerRepository,
		serviceRepository,
		&services.FormatFloat{},
		&services.FormatInt{},
		exporterContext,
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
