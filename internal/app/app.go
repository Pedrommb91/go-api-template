package app

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	server := api.NewServer(cfg, l)
	server.ServerConfigure()
	server.SetRoutes()
	server.Run()
}
