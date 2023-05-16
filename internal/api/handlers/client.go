package handlers

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"gorm.io/gorm"
)

type client struct {
	cfg *config.Config
	log logger.Interface
	db  *gorm.DB
}

func NewClient(cfg *config.Config, l logger.Interface, db *gorm.DB) openapi.ServerInterface {
	return &client{
		cfg: cfg,
		log: l,
		db:  db,
	}
}
