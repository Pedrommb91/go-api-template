//go:generate oapi-codegen -old-config-style -generate gin -o openapi/openapi_gin.gen.go -package openapi ../../spec/openapi.yaml
//go:generate oapi-codegen -old-config-style -generate spec -o openapi/openapi_spec.gen.go -package openapi ../../spec/openapi.yaml
//go:generate oapi-codegen -old-config-style -generate types -o openapi/openapi_types.gen.go -package openapi ../../spec/openapi.yaml
//go:generate oapi-codegen -old-config-style -generate client -o openapi/openapi_client.gen.go -package openapi ../../spec/openapi.yaml

package api

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/repositories"
	"github.com/Pedrommb91/go-api-template/internal/router"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *config.Config
	log    *logger.Logger
	engine *gin.Engine
	server *http.Server
	db     *repositories.PostgresDB
}

func NewServer(c *config.Config, l *logger.Logger, db *repositories.PostgresDB) *Server {
	handler := gin.New()
	return &Server{
		cfg:    c,
		log:    l,
		engine: handler,
		server: &http.Server{
			Addr:              c.API.Address,
			Handler:           handler,
			ReadHeaderTimeout: time.Second * 30,
		},
		db: db,
	}
}

func (s *Server) ServerConfigure() {
	if len(s.cfg.API.CORSAllowOrigins) > 0 {
		conf := cors.DefaultConfig()
		conf.AllowOrigins = s.cfg.API.CORSAllowOrigins
		conf.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
		conf.AllowCredentials = true
		s.engine.Use(cors.New(conf))
	}

	s.engine.NoRoute(func(ctx *gin.Context) {
		s.log.Error("Route not found")
	})
}

func (s *Server) SetRoutes() {
	router.NewRouter(s.engine, s.log, s.cfg, s.db)
}

func (s *Server) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func(logger.Interface) {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Fatal("listen: %s\n", err)
		}
	}(s.log)

	s.log.Info("%s started", s.cfg.App.Name)
	s.log.Info("version: %s", s.cfg.App.Version)

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	s.log.Info("App shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.log.Error("App server forced to shutdown: ", err)
	}

	s.log.Info("App exiting")
}
