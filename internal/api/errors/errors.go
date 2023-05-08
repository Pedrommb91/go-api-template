package errors

import (
	"net/http"

	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/clock"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/gin-gonic/gin"
)

//go:generate mockery --name ErrorSender
type ErrorSender interface {
	AbortWithStatusJSON(ctx *gin.Context, err error)
	JSON(ctx *gin.Context, err error)
	JSONWithStatus(ctx *gin.Context, err error, status int)
}

type APIError struct {
	clock clock.Clock
}

func NewErrorSender(clock clock.Clock) *APIError {
	return &APIError{
		clock: clock,
	}
}

func (ae *APIError) AbortWithStatusJSON(ctx *gin.Context, err error) {
	e := errors.GetFirstNestedError(err)
	er := newError(ctx, e, ae.clock)
	ctx.AbortWithStatusJSON(int(er.Status), er)
}

func (ae *APIError) JSON(ctx *gin.Context, err error) {
	e := errors.GetFirstNestedError(err)
	er := newError(ctx, e, ae.clock)
	ctx.JSON(int(er.Status), er)
}

func (ae *APIError) JSONWithStatus(ctx *gin.Context, err error, status int) {
	e := errors.GetFirstNestedError(err)
	er := newError(ctx, e, ae.clock)
	ctx.JSON(status, er)
}

func newError(ctx *gin.Context, err error, clock clock.Clock) *openapi.Error {
	er, ok := err.(*errors.Error)
	if !ok {
		return unexpectedError(ctx, clock)
	}
	return &openapi.Error{
		Id:        er.ID,
		Error:     er.Kind.String(),
		Message:   er.Message,
		Path:      ctx.FullPath(),
		Status:    int32(er.Kind.Int()),
		Timestamp: clock.Now(),
	}
}

func unexpectedError(ctx *gin.Context, clock clock.Clock) *openapi.Error {
	return &openapi.Error{
		Error:   errors.Unexpected.String(),
		Message: "Unexpected error occurred",
		Path:    ctx.FullPath(),
		Status:  http.StatusInternalServerError,
	}
}
