package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/config"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/handlers"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"
)

func main() {
	cfg := config.LoadConfig()

	services.InitMneeService(cfg)

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/balance/:address", handlers.GetBalance)
		api.GET("/balance", handlers.GetBalances)

		api.GET("/utxos/paginated/:address", handlers.GetPaginatedUtxos)
		api.GET("/utxos/:address/all", handlers.GetAllUtxos)

		api.GET("/transaction/:address", handlers.GetHistory)

		api.POST("/transaction/transfer", handlers.TransferSync)
		api.POST("/transaction/transfer-async", handlers.TransferAsync)
		api.POST("/transaction/submit-rawtx", handlers.SubmitRawTxAsync)
	}

	log.Printf("MNEE API Wrapper running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
