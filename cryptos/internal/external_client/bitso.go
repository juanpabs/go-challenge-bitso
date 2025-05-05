package externalclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const bitsoBaseURL = "https://api.bitso.com/v3"

type BookName string

const (
	BTC_MXN BookName = "btc_mxn"
	ETH_MXN BookName = "eth_mxn"
	XRP_MXN BookName = "xrp_mxn"
	BTC_USD BookName = "btc_usd"
	ETH_USD BookName = "eth_usd"
	XRP_USD BookName = "xrp_usd"
)

type BitsoError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BitsoTickerResponse struct {
	Error   BitsoError `json:"error,omitempty"`
	Success bool       `json:"success"`
	Payload struct {
		Book      string    `json:"book"`
		Volume    string    `json:"volume"`
		High      string    `json:"high"`
		Last      string    `json:"last"`
		Low       string    `json:"low"`
		Vwap      string    `json:"vwap"`
		Ask       string    `json:"ask"`
		Bid       string    `json:"bid"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"payload"`
}

func NewBitsoClient() Client {
	return &bitsoClient{baseUrl: bitsoBaseURL}
}

type bitsoClient struct {
	baseUrl string // Base URL for the Bitso API
}

type Client interface {
	GetTicker(book BookName) (BitsoTickerResponse, error)
}

func (c *bitsoClient) GetTicker(book BookName) (t BitsoTickerResponse, err error) {
	url := fmt.Sprintf("%s/ticker?book=%s", c.baseUrl, book)
	resp, err := http.Get(url)
	if err != nil {
		return t, fmt.Errorf("error trying to reach bitso api: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return t, fmt.Errorf("bitso api returns: %d", resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return t, fmt.Errorf("failed to decode response: %w", err)
	}

	if !t.Success {
		return t, fmt.Errorf("bitso API error response: (%v)%s", t.Error.Code, t.Error.Message)
	}

	return t, nil
}
