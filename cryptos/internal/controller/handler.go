package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanpabs/go-challenge-bitso/cryptos/internal/domain"
)

type CryptoUseCase interface {
	GetAllCryptos() ([]domain.Crypto, error)
}

type CryptoHandler interface {
	GetCryptos(*gin.Context)
}

type cryptoHandler struct {
	cryptoUseCase CryptoUseCase
}

func NewCryptoHandler(cryptoUseCase CryptoUseCase) CryptoHandler {
	return &cryptoHandler{
		cryptoUseCase: cryptoUseCase,
	}
}

func (h *cryptoHandler) GetCryptos(c *gin.Context) {
	cryptos, err := h.cryptoUseCase.GetAllCryptos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los precios"})
		return
	}
	c.JSON(http.StatusOK, cryptos)
}
