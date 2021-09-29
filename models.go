package main

type Account struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

type Deposit struct {
	Destination Account `json:"destination"`
}

type Withdraw struct {
	Origin Account `json:"origin"`
}

type Transfer struct {
	Origin      Account `json:"origin"`
	Destination Account `json:"destination"`
}

type Event struct {
	Type        string  `json:"type"`
	Destination string  `json:"destination,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Origin      string  `json:"origin,omitempty"`
}
