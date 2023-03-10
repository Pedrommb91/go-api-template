package service

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
)

type client struct {
	cfg *config.Config
	log logger.Interface
}

func NewClient(cfg *config.Config, l logger.Interface) openapi.ServerInterface {
	return &client{
		cfg: cfg,
		log: l,
	}
}
