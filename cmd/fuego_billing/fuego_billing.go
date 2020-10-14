package main

import (
	"fmt"
	"os"

	"fuegobyp-billing.com/internal/adapter"
	"fuegobyp-billing.com/internal/repository"
)

func main() {
	accountAdapter := adapter.NewAccountJSONAdapter()
	accountRepository := repository.NewAccountRepository(accountAdapter)

	customerAdapter := adapter.NewCustomerJSONAdapter()
	customerRepository := repository.NewCustomerRepository(customerAdapter)

	cmd, err := newRootCmd(os.Stdout, accountRepository, customerRepository)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}

	// @todo:

	// serviceTVA := model.TVA{Pourcent: 20}
	// service := model.Service{Detail: "Prestation de développement Bivwak BNP Paribas", Quantity: 20, UnitPrice: 663, TVA: serviceTVA}

}
