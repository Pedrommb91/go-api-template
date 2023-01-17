package app

import (
	"github.com/Pedrommb91/go-api-template/internal/api"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
)

func Run() {
	l := logger.New("info")

	server := api.NewServer(*l)
	server.ServerConfigure()
	server.SetRoutes()
	server.Run()
}
