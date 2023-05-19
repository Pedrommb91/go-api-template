package handlers

import (
	"net/http"

	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/gin-gonic/gin"
)

// GetUsersHandler implements openapi.ServerInterface
func (h *client) GetUsersHandler(c *gin.Context) {
	const op errors.Op = "handlers.GetUsersHandler"

	users, err := h.db.GetUsers()
	if err != nil {
		c.Error(errors.Build(
			errors.WithOp(op),
			errors.WithError(err),
			errors.WithMessage("Failed to get users"),
		))
		return
	}

	c.JSON(http.StatusOK, users)
}
