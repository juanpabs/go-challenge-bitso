package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/juanpabs/go-challenge-bitso/cryptos/internal/domain"
)

type BitsoTickerResponse struct {
	Success bool `json:"success"`
	Payload struct {
		Book      string `json:"book"`
		Volume    string `json:"volume"`
		High      string `json:"high"`
		Last      string `json:"last"`
		Low       string `json:"low"`
		Vwap      string `json:"vwap"`
		Ask       string `json:"ask"`
		Bid       string `json:"bid"`
		CreatedAt string `json:"created_at"`
	} `json:"payload"`
}

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
	// resp, err := http.Get("https://api.bitso.com/v3/ticker/?book=btc_mxn")
	// if err != nil {
	// 	log.Println("Error al hacer request a Bitso:", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el precio"})
	// 	return
	// }
	// defer resp.Body.Close()

	// var bitsoResp BitsoTickerResponse
	// if err := json.NewDecoder(resp.Body).Decode(&bitsoResp); err != nil {
	// 	log.Println("Error al decodificar respuesta de Bitso:", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Respuesta inválida de Bitso"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"last_price": bitsoResp.Payload.Last,
	// 	"high":       bitsoResp.Payload.High,
	// 	"low":        bitsoResp.Payload.Low,
	// 	"volume":     bitsoResp.Payload.Volume,
	// })
	c.JSON(200, nil)
}
