package router

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/handlers"
	"github.com/Pedrommb91/go-api-template/internal/api/middlewares"
	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/clock"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewRouter(engine *gin.Engine, l logger.Interface, cfg *config.Config) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.Use(middlewares.ErrorHandler(&clock.RealClock{}, l))

	mid := make([]openapi.MiddlewareFunc, 0)
	opt := openapi.GinServerOptions{
		BaseURL:     "/api/v1/",
		Middlewares: mid,
	}
	openapi.RegisterHandlersWithOptions(engine, handlers.NewClient(cfg, l), opt)
}
