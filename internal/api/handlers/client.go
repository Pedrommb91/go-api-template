package handlers

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/database"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
)

type client struct {
	cfg *config.Config
	log logger.Interface
	db  *database.PostgresDB
}

func NewClient(cfg *config.Config, l logger.Interface, db *database.PostgresDB) openapi.ServerInterface {
	return &client{
		cfg: cfg,
		log: l,
		db:  db,
	}
}
