package api

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Server struct {
	log    logger.Logger
	engine *gin.Engine
	server *http.Server
}

func NewServer(l logger.Logger) *Server {
	handler := gin.New()
	return &Server{
		log:    l,
		engine: handler,
		server: &http.Server{
			Addr:              ":8080",
			Handler:           handler,
			ReadHeaderTimeout: time.Second * 30,
		},
	}
}

func (s *Server) ServerConfigure() {
	s.engine.NoRoute(func(ctx *gin.Context) {
		s.log.Error("Route not found")
	})
}

func (s *Server) SetRoutes() {

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
	}(&s.log)

	s.log.Info("App started")

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
