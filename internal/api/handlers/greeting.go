package handlers

import (
	"net/http"
	"time"

	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/gin-gonic/gin"
)

// (GET /greeting/{name})
func (g *client) GetHelloWorldHandler(c *gin.Context, name string) {
	if name == "D. Sebastian" {
		msg := "D. Sebastian did not appear yet"
		g.log.Error(msg)
		c.JSON(http.StatusInternalServerError, &openapi.Error{
			Error:     "Not returned",
			Message:   msg,
			Path:      "/greeting/" + name,
			Status:    500,
			Timestamp: time.Now(),
		})
		return
	}

	c.JSON(http.StatusOK, &openapi.Greeting{
		Message: "Hello " + name,
	})
}
