package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type Base struct {
// 	ID        string `json:"id"`
// 	CreatedAt time.Timer
// 	UpdatedAt time.Timer
// 	DeletedAt time.Timer
// }

var Accounts []Account

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/reset", ResetStatus).Methods("Post")
	router.HandleFunc("/balance", GetBalance).Methods("Get")
	router.HandleFunc("/event", PostEvent).Methods("Post")

	Accounts = append(Accounts, Account{ID: "300", Balance: 0})
	log.Fatal(http.ListenAndServe(":5050", router))
}
