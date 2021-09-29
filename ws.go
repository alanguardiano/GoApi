package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ResetStatus(w http.ResponseWriter, r *http.Request) {
	// # Reset state before starting tests
	// POST /reset
	// 200 OK

	Accounts = nil
	Accounts = append(Accounts, Account{ID: "300", Balance: 0})

	w.WriteHeader(http.StatusOK)
	respOK := strings.Trim(`"OK"`, "\"")

	fmt.Fprintf(w, respOK)
	return

}
func GetBalance(w http.ResponseWriter, r *http.Request) {
	// # Get balance for existing account
	// GET /balance?account_id=100
	// 200 20

	// # Get balance for non-existing account
	// GET /balance?account_id=1234
	// 404 0

	account_id := r.URL.Query().Get("account_id")

	if account_id == "" {
		json.NewEncoder(w).Encode("account not informed")
		return
	}

	fmt.Println(Accounts)
	for _, account := range Accounts {
		if account.ID == account_id {
			json.NewEncoder(w).Encode(account.Balance)
			return
		}
	}

	//json.NewEncoder(w).Encode("None Accounts Registered")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(0)
	return

}
func PostEvent(w http.ResponseWriter, r *http.Request) {
	// # Create account with initial balance
	// POST /event {"type":"deposit", "destination":"100", "amount":10}
	// 201 {"destination": {"id":"100", "balance":10}}
	var event Event

	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if event.Type != "" {
		if event.Type != "deposit" && event.Type != "transfer" && event.Type != "withdraw" {
			fmt.Println("EVENT NOT REGISTERED")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(0)
			return
		}
	} else {
		fmt.Println("TYPE not informed on json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(0)
		return
	}

	switch event.Type {
	case "deposit":
		fmt.Println(Accounts)
		account, err := DepositEvent(event)
		fmt.Println(Accounts)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(0)
			return
		}
		w.WriteHeader(http.StatusCreated)
		// fmt.Fprintf(w, "Success")
		var deposit Deposit
		deposit.Destination = account
		json.NewEncoder(w).Encode(deposit)
		return
	case "withdraw":
		fmt.Println("withdraw - CHAMAR FUNCAO withdraw")
		acc, err := WithdrawEvent(event)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			//fmt.Fprintf(w, err.Error())
			json.NewEncoder(w).Encode(0)
			return
		}
		w.WriteHeader(http.StatusCreated)
		var withdraw Withdraw
		withdraw.Origin = acc
		json.NewEncoder(w).Encode(withdraw)
		return
	case "transfer":
		fmt.Println("transfer - CHAMAR FUNCAO transfer")
		// idxOrigin, idxDestination, err := Transfer(event)
		idxOrigin, idxDestination, err := TransferEvent(event)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			// fmt.Fprintf(w, err.Error())
			json.NewEncoder(w).Encode(0)
			return
		}
		w.WriteHeader(http.StatusCreated)
		var transfer Transfer
		transfer.Origin = Accounts[idxOrigin]
		transfer.Destination = Accounts[idxDestination]
		json.NewEncoder(w).Encode(transfer)
		return
	}

	// fmt.Fprintf(w, "Event: %+v", event)
}
