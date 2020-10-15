package main

import (
	"fmt"
	"os"

	"fuegobyp-billing.com/internal/adapter"
	"fuegobyp-billing.com/internal/repository"
	"fuegobyp-billing.com/internal/services"
)

func main() {
	accountAdapter := adapter.NewAccountJSONAdapter()
	accountRepository := repository.NewAccountRepository(accountAdapter)

	customerAdapter := adapter.NewCustomerJSONAdapter()
	customerRepository := repository.NewCustomerRepository(customerAdapter)

	serviceAdapter := adapter.NewServiceJSONAdapter()
	serviceRepository := repository.NewServiceRepository(serviceAdapter)

	cmd, err := newRootCmd(os.Stdout, accountRepository, customerRepository, serviceRepository, &services.FormatFloat{}, &services.FormatInt{})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
