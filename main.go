package main

import (
	"fmt"
	"os"

	"github.com/JamshedJ/loanly/config"
	adminV1 "github.com/JamshedJ/loanly/delivery/rest/admin/v1"
	loanV1 "github.com/JamshedJ/loanly/delivery/rest/loan/v1"
	"github.com/JamshedJ/loanly/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg := config.InitConfig("./config")

	svc := services.ServiceFacade{}

	adminApi := adminV1.AdminApiV1{
		Logger: logger.With().Str("component", "admin_api_v1").Logger(),
		Svc:    svc,
	}

	loanApi := loanV1.LoanApiV1{
		Logger: logger.With().Str("component", "loan_api_v1").Logger(),
		Svc:    svc,
	}

	api := gin.New()

	err := adminApi.RegisterRoutes(api)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to register admin api routes")
	}

	err = loanApi.RegisterRoutes(api)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to register loan api routes")
	}

	api.Run(fmt.Sprintf(":%d", cfg.App.Port))
}
