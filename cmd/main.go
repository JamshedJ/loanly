package main

import (
	"context"
	"os"

	restapi "github.com/JamshedJ/REST-api"
	"github.com/JamshedJ/REST-api/configs"
	"github.com/JamshedJ/REST-api/internal/handler"
	"github.com/JamshedJ/REST-api/internal/repository"
	"github.com/JamshedJ/REST-api/internal/service"
	"github.com/JamshedJ/REST-api/pkg/glog"
)

func main() {
	logger := glog.NewLogger()
	ctx := context.Background()

	dsn, err := configs.InitConfig()
	if err != nil {
		logger.Fatal().Ctx(ctx).Err(err).Msg("error initializing configuration")
	}

	db, err := repository.DBConnection(dsn)
	if err != nil {
		logger.Fatal().Ctx(ctx).Err(err).Msg("failed to initialize DBConnection")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(*repos)
	handlers := handler.NewHandler(*services)

	srv := new(restapi.Server)
	go func()  {
		if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
			logger.Fatal().Ctx(ctx).Err(err).Msg("failed running server")
		}
	}()

	quite := make(chan os.Signal, 1)
	os.Signal.Signal(<-quite)
	
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Ctx(ctx).Err(err).Msg("error occured on server shutting down")
	}
}
