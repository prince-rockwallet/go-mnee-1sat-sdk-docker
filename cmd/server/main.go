package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/config"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/handlers"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/mnee-xyz/go-mnee-1sat-sdk-docker/docs"
)

// @title           MNEE SDK API Wrapper (Go)
// @version         1.0
// @description     A high-performance Golang REST API wrapper for the MNEE Golang SDK.

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cfg := config.LoadConfig()

	services.InitMneeService(cfg)

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/balance/:address", handlers.GetBalance)
		api.GET("/balance", handlers.GetBalances)

		api.GET("/utxos/paginated", handlers.GetPaginatedUtxos)
		api.GET("/utxos/all", handlers.GetAllUtxos)

		api.GET("/transaction", handlers.GetHistory)

		api.POST("/transaction/transfer", handlers.TransferSync)
		api.POST("/transaction/transfer-async", handlers.TransferAsync)
		api.POST("/transaction/submit-rawtx", handlers.SubmitRawTxAsync)
	}

	r.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("MNEE API Wrapper running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
