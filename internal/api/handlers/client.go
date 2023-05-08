package handlers

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/errors"
	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
)

type client struct {
	cfg         *config.Config
	log         logger.Interface
	errorSender errors.ErrorSender
}

func NewClient(cfg *config.Config, l logger.Interface, es errors.ErrorSender) openapi.ServerInterface {
	return &client{
		cfg:         cfg,
		log:         l,
		errorSender: es,
	}
}
