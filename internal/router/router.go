package router

import (
	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/handlers"
	"github.com/Pedrommb91/go-api-template/internal/api/middlewares"
	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/internal/api/repositories"
	"github.com/Pedrommb91/go-api-template/pkg/clock"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

func NewRouter(engine *gin.Engine, l logger.Interface, cfg *config.Config, db *repositories.PostgresDB) {
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	engine.Use(middlewares.ErrorHandler(&clock.RealClock{}, l))

	// Swagger
	engine.StaticFile("/swagger", "./spec/openapi.yaml")

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger", Path: "/swagger-ui"}
	sh := middleware.SwaggerUI(opts, nil)
	engine.GET("/swagger-ui", func(ctx *gin.Context) {
		sh.ServeHTTP(ctx.Writer, ctx.Request)
	})

	mid := make([]openapi.MiddlewareFunc, 0)
	opt := openapi.GinServerOptions{
		BaseURL:     "/api/v1/",
		Middlewares: mid,
	}
	openapi.RegisterHandlersWithOptions(engine, handlers.NewClient(cfg, l, db), opt)
}
