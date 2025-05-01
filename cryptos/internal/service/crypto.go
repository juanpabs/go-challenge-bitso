package service

type CryptoService interface {
	GetCryptoValue()
}

type cryptoService struct{}

func NewCryptoService() CryptoService {
	return &cryptoService{}
}

func (s *cryptoService) GetCryptoValue() {}
