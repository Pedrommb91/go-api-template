// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/gin-gonic/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /greeting/{name})
	GetHelloWorld(c *gin.Context, name string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(c *gin.Context)

// GetHelloWorld operation middleware
func (siw *ServerInterfaceWrapper) GetHelloWorld(c *gin.Context) {

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameter("simple", false, "name", c.Param("name"), &name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Invalid format for parameter name: %s", err)})
		return
	}

	c.Set(BasicAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetHelloWorld(c, name)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	router.GET(options.BaseURL+"/greeting/:name", wrapper.GetHelloWorld)

	return router
}