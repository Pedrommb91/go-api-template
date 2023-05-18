package handlers

import (
	"fmt"
	"net/http"

	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/gin-gonic/gin"
)

// (GET /greeting/{name})
func (g *client) GetHelloWorldHandler(c *gin.Context, name string) {
	if name == "D. Sebastian" {
		msg := "D. Sebastian did not appear yet"

		//nolint:errcheck
		c.Error(errors.Build(
			errors.WithError(fmt.Errorf("not returned")),
			errors.WithMessage(msg),
			errors.KindNotFound(),
		))
		return
	}

	c.JSON(http.StatusOK, &openapi.Greeting{
		Message: "Hello " + name,
	})
}
