package api

import (
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/repositories"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

func TestServer_Run(t *testing.T) {
	handler := gin.New()
	type fields struct {
		cfg    *config.Config
		log    *logger.Logger
		engine *gin.Engine
		server *http.Server
	}
	type args struct {
		db *repositories.PostgresDB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "valid args runs ok",
			fields: fields{
				cfg: &config.Config{
					API: config.API{
						Address:          ":0",
						CORSAllowOrigins: []string{"https://ttt.com"},
					},
				},
				log:    logger.New("info"),
				engine: handler,
				server: &http.Server{
					Addr:              ":0",
					Handler:           handler,
					ReadHeaderTimeout: time.Second * 30,
				},
			},
			args: args{
				db: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewServer(tt.fields.cfg, tt.fields.log, tt.args.db)
			s.ServerConfigure()
			s.SetRoutes()
			go func() {
				time.Sleep(time.Millisecond * 500)
				p, err := os.FindProcess(os.Getpid())
				if err != nil {
					t.Errorf("Could not find process")
				}
				err = p.Signal(syscall.SIGINT)
				if err != nil {
					t.Errorf("Could not find process")
				}
			}()
			s.Run()
		})
	}
}
