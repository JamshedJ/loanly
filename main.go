package main

import (
	"fmt"
	"os"

	"github.com/JamshedJ/loanly/config"
	v1 "github.com/JamshedJ/loanly/delivery/rest/admin/v1"
	"github.com/JamshedJ/loanly/domain/services"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg := config.InitConfig("./config")

	svc := services.ServiceFacade{}

	adminApi := v1.AdminApiV1{
		Logger: logger,
		Svc:    svc,
	}

	adminApi.RegisterRoutes().Run(fmt.Sprintf(":%d", cfg.App.Port))
}
