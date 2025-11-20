package services

import (
	"log"
	"strings"

	mnee "github.com/mnee-xyz/go-mnee-1sat-sdk"
	"github.com/mnee-xyz/go-mnee-1sat-sdk-docker/internal/config"
)

var Instance *mnee.MNEE

func InitMneeService(cfg *config.Config) {
	if cfg.MneeApiKey == "" {
		log.Fatal("MNEE_API_KEY is required in .env")
	}

	var targetEnv string
	if strings.ToLower(cfg.MneeEnv) == "production" {
		targetEnv = mnee.EnvMain
	} else {
		targetEnv = mnee.EnvSandbox
	}

	var err error
	Instance, err = mnee.NewMneeInstance(targetEnv, cfg.MneeApiKey)
	if err != nil {
		log.Fatalf("Failed to initialize MNEE SDK: %v", err)
	}

	log.Printf("MNEE SDK Initialized in %s mode", targetEnv)
}
