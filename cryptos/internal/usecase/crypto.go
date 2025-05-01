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
	GetCryptoValue()
}

func NewCryptoUseCase(srv CryptoService) CryptoUseCase {
	return &cryptoUseCase{
		cryptoService: srv,
	}
}

func (uc *cryptoUseCase) GetAllCryptos() ([]domain.Crypto, error) {
	return nil, nil
}
