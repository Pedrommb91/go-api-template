package app

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api"
	"github.com/Pedrommb91/go-api-template/pkg/database"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	db, err := database.OpenPostgres(cfg.Database)
	if err != nil {
		panic(err)
	}

	server := api.NewServer(cfg, l)
	server.ServerConfigure()
	server.SetRoutes(db)
	server.Run()
}
