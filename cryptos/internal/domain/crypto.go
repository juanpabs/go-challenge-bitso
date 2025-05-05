package domain

import (
	"strings"
	"time"

	externalclient "github.com/juanpabs/go-challenge-bitso/cryptos/internal/external_client"
)

type Price struct {
	USD string `json:"usd"`
	MXN string `json:"mxn"`
}

type CryptoCurrency string

const BTC CryptoCurrency = "BTC"
const ETH CryptoCurrency = "ETH"
const XRP CryptoCurrency = "XRP"

var AvailableCryptoCurrencies = []CryptoCurrency{
	BTC,
	ETH,
	XRP,
}

type Crypto struct {
	Id        int         `json:"id"`
	Component string      `json:"component"`
	Model     CryptoModel `json:"model"`
}

type CryptoModel struct {
	Date         time.Time `json:"date"`          // The time at which the price is fetched.
	Name         string    `json:"name"`          // The name of the Modelcurrency (e.g., Bitcoin).
	TickerSymbol string    `json:"ticker_symbol"` // The ticker symbol (e.g., BTC).
	Price        Price     `json:"price"`         // The price in different currencies.
}

func (c *Crypto) SetPrice(bookName externalclient.BookName, price string) {
	priceCurrency := strings.Split(string(bookName), "_")
	switch priceCurrency[1] {
	case "mxn":
		c.Model.Price.MXN = price
	case "usd":
		c.Model.Price.USD = price
	}
}

func (c *Crypto) SetTickerSymbol(bookName externalclient.BookName) {
	cryptoSymbol := strings.Split(string(bookName), "_")
	switch cryptoSymbol[0] {
	case "btc":
		c.Model.TickerSymbol = string(BTC)
	case "eth":
		c.Model.TickerSymbol = string(ETH)
	case "xrp":
		c.Model.TickerSymbol = string(XRP)
	}
}

func (c *Crypto) SetName(bookName externalclient.BookName) {
	name := strings.Split(string(bookName), "_")
	switch name[0] {
	case "btc":
		c.Model.Name = "Bitcoin"
	case "eth":
		c.Model.Name = "Ethereum"
	case "xrp":
		c.Model.Name = "XRP"
	}
}

func (c *Crypto) SetId(id int) {
	c.Id = id
}

func (c *Crypto) SetComponent(bookName externalclient.BookName) {
	cryptoSymbol := strings.Split(string(bookName), "_")
	c.Component = "component_" + cryptoSymbol[0]
}
