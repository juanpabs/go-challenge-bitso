package usecase

import (
	"github.com/juanpabs/go-challenge-bitso/cryptos/internal/domain"
	"github.com/juanpabs/go-challenge-bitso/cryptos/internal/service"
)

type CryptoUseCase interface {
	GetAllCryptos() ([]domain.Crypto, error)
}

type cryptoUseCase struct {
	cryptoService service.CryptoService
}

type CryptoService interface {
	GetCryptoValue(domain.CryptoCurrency) (domain.Crypto, error)
}

func NewCryptoUseCase(srv CryptoService) CryptoUseCase {
	return &cryptoUseCase{
		cryptoService: srv,
	}
}

func (uc *cryptoUseCase) GetAllCryptos() ([]domain.Crypto, error) {
	layout := []domain.Crypto{}
	for id, cryptoCurrency := range domain.AvailableCryptoCurrencies {
		crypto, err := uc.cryptoService.GetCryptoValue(cryptoCurrency)
		if err != nil {
			return nil, err
		}
		crypto.Id = id
		layout = append(layout, crypto)
	}
	return nil, nil
}
