package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/juanpabs/go-challenge-bitso/cryptos/internal/domain"
)

type cryptoUseCase interface {
	GetAllCryptos() ([]domain.Crypto, error)
}

func NewRouter(cryptoUsecase cryptoUseCase) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	cryptoHandler := NewCryptoHandler(cryptoUsecase)

	r.GET("/cryptos", cryptoHandler.GetCryptos)

	return r
}
