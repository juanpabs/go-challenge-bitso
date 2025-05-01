package main

import (
	"github.com/juanpabs/go-challenge-bitso/cryptos/internal/controller"
	"github.com/juanpabs/go-challenge-bitso/cryptos/internal/service"
	"github.com/juanpabs/go-challenge-bitso/cryptos/internal/usecase"
)

func main() {

	cryptoService := service.NewCryptoService()
	cryptoUseCase := usecase.NewCryptoUseCase(cryptoService)
	r := controller.NewRouter(cryptoUseCase)
	r.Run(":8080")
}
