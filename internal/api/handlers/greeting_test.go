package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Pedrommb91/go-api-template/config"
	"github.com/Pedrommb91/go-api-template/internal/api/middlewares"
	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/internal/api/repositories"
	"github.com/Pedrommb91/go-api-template/pkg/clock/mocks"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_client_GetHelloWorldHandler(t *testing.T) {
	now := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	dummyErrID := "e157f89f-abd0-4b1a-bc58-de8bd8fd04cd"
	errors.NewUUID = func() uuid.UUID { return uuid.FromStringOrNil(dummyErrID) }

	type fields struct {
		cfg *config.Config
		log logger.Interface
	}
	type args struct {
		name string
		db   *repositories.PostgresDB
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantCode         int
		expectedResponse *openapi.Greeting
		expectedErr      *openapi.Error
	}{
		{
			name: "user greeting",
			fields: fields{
				cfg: &config.Config{},
				log: logger.New("info"),
			},
			args: args{
				name: "user",
				db:   nil,
			},
			wantCode: http.StatusOK,
			expectedResponse: &openapi.Greeting{
				Message: "Hello user",
			},
			expectedErr: nil,
		},
		{
			name: "D. Sebastian greeting not found",
			fields: fields{
				cfg: &config.Config{},
				log: logger.New("info"),
			},
			args: args{
				name: "D. Sebastian",
				db:   nil,
			},
			wantCode:         http.StatusNotFound,
			expectedResponse: nil,
			expectedErr: &openapi.Error{
				Error:     "Not Found",
				Id:        dummyErrID,
				Message:   "D. Sebastian did not appear yet",
				Path:      "/api/v1/greeting/D. Sebastian",
				Status:    http.StatusNotFound,
				Timestamp: now,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clock := mocks.NewClock(t)
			clock.On("Now").Return(now).Maybe()

			r := gin.Default()
			g := NewClient(tt.fields.cfg, tt.fields.log, tt.args.db)
			r.Use(middlewares.ErrorHandler(clock, tt.fields.log))
			r.GET("/api/v1/greeting/"+tt.args.name, func(c *gin.Context) {
				g.GetHelloWorldHandler(c, tt.args.name)
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/greeting/"+tt.args.name, nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wantCode, w.Code)

			if w.Code == http.StatusOK {
				response := &openapi.Greeting{}
				if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
					t.Errorf("Failed to decode message: %d", err)
					t.FailNow()
				}

				assert.Equal(t, tt.expectedResponse, response)
			} else {
				response := &openapi.Error{}
				if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
					t.Errorf("Failed to decode message: %d", err)
					t.FailNow()
				}
				assert.Equal(t, tt.expectedErr, response)
			}
		})
	}
}
