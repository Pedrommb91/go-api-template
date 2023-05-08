package router

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/errors"
	"github.com/Pedrommb91/go-api-template/internal/api/handlers"
	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(engine *gin.Engine, l logger.Interface, cfg *config.Config, es errors.ErrorSender) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	mid := make([]openapi.MiddlewareFunc, 0)
	opt := openapi.GinServerOptions{
		BaseURL:     "/api/v1/",
		Middlewares: mid,
	}
	openapi.RegisterHandlersWithOptions(engine, handlers.NewClient(cfg, l, es), opt)
}
