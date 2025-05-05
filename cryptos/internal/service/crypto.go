package service

import (
	"fmt"

	"github.com/juanpabs/go-challenge-bitso/cryptos/internal/domain"
	externalclient "github.com/juanpabs/go-challenge-bitso/cryptos/internal/external_client"
)

type CryptoService interface {
	GetCryptoValue(domain.CryptoCurrency) (domain.Crypto, error)
}

type cryptoService struct{}

func NewCryptoService() CryptoService {
	return &cryptoService{}
}

func (s *cryptoService) GetCryptoValue(crypto domain.CryptoCurrency) (domain.Crypto, error) {
	return GetCryptoFromBitso(crypto)
}

func GetBookNamesFromCrypto(crypto domain.CryptoCurrency) ([]externalclient.BookName, error) {
	switch crypto {
	case domain.BTC:
		return []externalclient.BookName{externalclient.BTC_MXN, externalclient.BTC_USD}, nil
	case domain.ETH:
		return []externalclient.BookName{externalclient.ETH_MXN, externalclient.ETH_USD}, nil
	case domain.XRP:
		return []externalclient.BookName{externalclient.XRP_MXN, externalclient.XRP_USD}, nil
	default:
		return nil, fmt.Errorf("unsupported crypto currency: %s", crypto)
	}
}

func GetCryptoFromBitso(crypto domain.CryptoCurrency) (c domain.Crypto, err error) {
	bitsoClient := externalclient.NewBitsoClient()

	bookNames, err := GetBookNamesFromCrypto(crypto)
	if err != nil {
		return c, fmt.Errorf("error getting book names for crypto %s: %w", c, err)
	}

	for _, bookName := range bookNames {
		bitsoTicker, err := bitsoClient.GetTicker(bookName)
		if err != nil {
			return c, fmt.Errorf("error getting ticker for book %s: %w", bookName, err)
		}

		c.SetPrice(bookName, bitsoTicker.Payload.Last)
		c.Model.Date = bitsoTicker.Payload.CreatedAt
		c.SetTickerSymbol(bookName)
		c.SetName(bookName)

	}
	return c, nil
}
