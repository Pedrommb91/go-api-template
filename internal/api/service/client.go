package service

import (
	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
)

type client struct {
	log logger.Interface
}

func NewClient(l logger.Interface) openapi.ServerInterface {
	return &client{
		log: l,
	}
}
