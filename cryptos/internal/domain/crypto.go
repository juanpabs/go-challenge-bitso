package domain

import "time"

type Price struct {
	USD string `json:"usd"`
	MXN string `json:"mxn"`
}

type Crypto struct {
	Date         time.Time `json:"date"`          // The time at which the price is fetched.
	Name         string    `json:"name"`          // The name of the cryptocurrency (e.g., Bitcoin).
	TickerSymbol string    `json:"ticker_symbol"` // The ticker symbol (e.g., BTC).
	Price        Price     `json:"price"`         // The price in different currencies.
}
