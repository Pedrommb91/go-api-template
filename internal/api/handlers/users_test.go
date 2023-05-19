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
	"github.com/Pedrommb91/go-api-template/pkg/clock/mocks"
	"github.com/Pedrommb91/go-api-template/pkg/database"
	"github.com/Pedrommb91/go-api-template/pkg/database/tests"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_client_GetUsersHandler(t *testing.T) {
	now := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	dummyErrID := "e157f89f-abd0-4b1a-bc58-de8bd8fd04cd"
	errors.NewUUID = func() uuid.UUID { return uuid.FromStringOrNil(dummyErrID) }

	db := tests.NewPostgresTestContainer()

	type fields struct {
		cfg *config.Config
		log logger.Interface
		db  *database.PostgresDB
	}
	tests := []struct {
		name                  string
		fields                fields
		wantCode              int
		expectedUsersResponse []*openapi.GetUsersResponse
		expectedErrorResponse *openapi.Error
	}{
		{
			name: "Valid args return users",
			fields: fields{
				cfg: &config.Config{},
				log: logger.New("info"),
				db:  db,
			},
			wantCode: http.StatusOK,
			expectedUsersResponse: []*openapi.GetUsersResponse{
				{
					Id:   "1",
					Name: "test",
				},
			},
			expectedErrorResponse: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clock := mocks.NewClock(t)
			clock.On("Now").Return(now).Maybe()

			r := gin.Default()
			r.Use(middlewares.ErrorHandler(clock, tt.fields.log))

			g := NewClient(tt.fields.cfg, tt.fields.log, tt.fields.db)
			r.GET("/users", func(c *gin.Context) {
				g.GetUsersHandler(c)
			})

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/users", nil)

			r.ServeHTTP(w, req)
			assert.Equal(t, tt.wantCode, w.Code)

			if tt.expectedUsersResponse != nil {
				var got []*openapi.GetUsersResponse
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Errorf("Failed to unmarshal body: %s", err)
				}
				assert.Equal(t, got, tt.expectedUsersResponse)
			}
			if tt.expectedErrorResponse != nil {
				var got *openapi.Error
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Errorf("Failed to unmarshal body: %s", err)
				}
				assert.Equal(t, got, tt.expectedErrorResponse)
			}
		})
	}
}
