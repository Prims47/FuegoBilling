package main

import (
	"fmt"
	"os"

	"github.com/prims47/FuegoBilling/internal/adapter"
	"github.com/prims47/FuegoBilling/internal/repository"
	"github.com/prims47/FuegoBilling/internal/services"
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
