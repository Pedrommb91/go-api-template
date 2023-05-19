package app

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api"
	"github.com/Pedrommb91/go-api-template/internal/api/repositories"
	"github.com/Pedrommb91/go-api-template/pkg/database"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	db := &repositories.PostgresDB{DB: database.NewPostgresOrDie(cfg.Database)}

	server := api.NewServer(cfg, l, db)
	server.ServerConfigure()
	server.SetRoutes()
	server.Run()
}
