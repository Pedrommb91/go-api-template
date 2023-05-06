package errors

import (
	"net/http"
	"time"

	"github.com/Pedrommb91/go-api-template/internal/api/openapi"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/gin-gonic/gin"
)

var Now = time.Now

func NewError(ctx *gin.Context, err error) *openapi.Error {
	er, ok := err.(*errors.Error)
	if !ok {
		return unexpectedError(ctx)
	}
	return &openapi.Error{
		Id:        er.ID,
		Error:     er.Kind.String(),
		Message:   er.Message,
		Path:      ctx.FullPath(),
		Status:    int32(er.Kind.Int()),
		Timestamp: Now(),
	}
}

func AbortWithStatusJSON(ctx *gin.Context, err error) {
	e := errors.GetFirstNestedError(err)
	er := NewError(ctx, e)
	ctx.AbortWithStatusJSON(int(er.Status), er)
}

func JSON(ctx *gin.Context, err error) {
	e := errors.GetFirstNestedError(err)
	er := NewError(ctx, e)
	ctx.JSON(int(er.Status), er)
}

func JSONWithStatus(ctx *gin.Context, err error, status int) {
	e := errors.GetFirstNestedError(err)
	er := NewError(ctx, e)
	ctx.JSON(status, er)
}

func unexpectedError(ctx *gin.Context) *openapi.Error {
	return &openapi.Error{
		Error:     errors.Unexpected.String(),
		Message:   "Unexpected error occurred",
		Path:      ctx.FullPath(),
		Status:    http.StatusInternalServerError,
		Timestamp: Now(),
	}
}
