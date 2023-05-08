package app

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api"
	"github.com/Pedrommb91/go-api-template/internal/api/errors"
	"github.com/Pedrommb91/go-api-template/pkg/clock"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	server := api.NewServer(cfg, l, errors.NewErrorSender(&clock.RealClock{}))
	server.ServerConfigure()
	server.SetRoutes()
	server.Run()
}
