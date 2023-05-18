package handlers

import (
	"database/sql"

	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
)

type client struct {
	cfg *config.Config
	log logger.Interface
	db  *sql.DB
}

func NewClient(cfg *config.Config, l logger.Interface, db *sql.DB) openapi.ServerInterface {
	return &client{
		cfg: cfg,
		log: l,
		db:  db,
	}
}
