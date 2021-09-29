package main

import (
	"errors"
	"fmt"
)

func DepositEvent(event Event) (account Account, err error) {
	// # Create account with initial balance
	// POST /event {"type":"deposit", "destination":"100", "amount":10}
	// 201 {"destination": {"id":"100", "balance":10}}

	// # Deposit into existing account
	// POST /event {"type":"deposit", "destination":"100", "amount":10}
	// 201 {"destination": {"id":"100", "balance":20}}

	if event.Amount < 0 {
		err := errors.New("amount can't be negative")
		return account, err
	}

	for i := 0; i < len(Accounts); i++ {
		if Accounts[i].ID == event.Destination {
			// acc := Accounts[i]
			// acc.Balance += event.Amount
			Accounts[i].Balance += event.Amount
			account = Accounts[i]
			return account, nil
		}
	}

	var acc Account
	acc.ID = event.Destination
	acc.Balance = event.Amount
	Accounts = append(Accounts, acc)
	account = acc

	return account, nil
}
func TransferEvent(event Event) (indexOrigin, indexDestination int, err error) {
	// # Transfer from existing account
	// POST /event {"type":"transfer", "origin":"100", "amount":15, "destination":"300"}
	// 201 {"origin": {"id":"100", "balance":0}, "destination": {"id":"300", "balance":15}}

	// # Transfer from non-existing account
	// POST /event {"type":"transfer", "origin":"200", "amount":15, "destination":"300"}
	// 404 0
	var (
		idxOrigin      int = -1
		idxDestination int = -1
	)

	if event.Amount < 0 {
		err := errors.New("amount can't be negative")
		return -1, -1, err
	}

	for i := 0; i < len(Accounts); i++ {
		if Accounts[i].ID == event.Origin {
			// fmt.Println("conta origem existe")
			idxOrigin = i
		}
	}
	for i := 0; i < len(Accounts); i++ {
		if Accounts[i].ID == event.Destination {
			// fmt.Println("conta destino existe")
			idxDestination = i
		}
	}

	if idxOrigin == -1 {
		err := errors.New("origin account do not exists")
		return -1, -1, err
	}

	if idxDestination == -1 {
		err := errors.New("destination account do not exists")
		return -1, -1, err
	}

	fmt.Println(Accounts)
	//VALIDAR SE O SALDO DA CONTA ORIGEM É NEGATIVO???
	//TRUE NAO DEIXA FAZER A TRANSFERENCIA
	//FALSE FAZ A TRANSFERENCIA
	//PS: nada na documentação então não implementado...

	Accounts[idxOrigin].Balance -= event.Amount
	Accounts[idxDestination].Balance += event.Amount

	fmt.Println(Accounts)
	return idxOrigin, idxDestination, nil
}

func WithdrawEvent(event Event) (account Account, err error) {
	// # Withdraw from non-existing account
	// POST /event {"type":"withdraw", "origin":"200", "amount":10}
	// 404 0

	// # Withdraw from existing account
	// POST /event {"type":"withdraw", "origin":"100", "amount":5}
	// 201 {"origin": {"id":"100", "balance":15}}
	var acc Account

	if event.Amount < 0 {
		err := errors.New("amount can't be negative")
		return account, err
	}

	for i := 0; i < len(Accounts); i++ {
		if Accounts[i].ID == event.Origin {
			fmt.Println("conta existe, fazer o saque")
			acc = Accounts[i]
			Accounts[i].Balance -= event.Amount
			account = Accounts[i]
			return
		}
	}

	if acc.ID == "" {
		err := errors.New("account don't exists")
		return account, err
	}

	return
}
