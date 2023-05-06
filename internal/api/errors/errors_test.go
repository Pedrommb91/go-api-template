package errors

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/gin-gonic/gin"
)

func TestAbortWithStatusJSON(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	type args struct {
		ctx *gin.Context
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Valid costum error, context is aborted and error returned",
			args: args{
				ctx: c,
				err: errors.Build(),
			},
		},
		{
			name: "Valid native error, context is aborted and error returned",
			args: args{
				ctx: c,
				err: fmt.Errorf("error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AbortWithStatusJSON(tt.args.ctx, tt.args.err)
			if !c.IsAborted() {
				t.Errorf("Context not aborted")
			}
			for _, err := range c.Errors {
				if err == nil {
					t.Errorf("Must contain error")
				}
				if !errors.Equal(errors.GetFirstNestedError(err), tt.args.err) {
					t.Errorf("AbortWithStatusJSON() got = %v, want %v", err, tt.args.err)
				}
			}
		})
	}
}

func TestJSON(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	type args struct {
		ctx *gin.Context
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Valid costum error, context is not aborted and error returned",
			args: args{
				ctx: c,
				err: errors.Build(),
			},
		},
		{
			name: "Valid native error, context is not aborted and error returned",
			args: args{
				ctx: c,
				err: fmt.Errorf("error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JSON(tt.args.ctx, tt.args.err)
			if c.IsAborted() {
				t.Errorf("Context aborted")
			}
			for _, err := range c.Errors {
				if err == nil {
					t.Errorf("Must contain error")
				}
				if !errors.Equal(errors.GetFirstNestedError(err), tt.args.err) {
					t.Errorf("JSON() got = %v, want %v", err, tt.args.err)
				}
			}
		})
	}
}

func TestJSONWithStatus(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	type args struct {
		ctx    *gin.Context
		err    error
		status int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Valid costum error, context is not aborted and error returned with status",
			args: args{
				ctx:    c,
				err:    errors.Build(),
				status: 400,
			},
		},
		{
			name: "Valid native error, context is not aborted and error returned",
			args: args{
				ctx:    c,
				err:    fmt.Errorf("error"),
				status: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JSONWithStatus(tt.args.ctx, tt.args.err, tt.args.status)
			if c.IsAborted() {
				t.Errorf("Context aborted")
			}
			if c.Writer.Status() != tt.args.status {
				t.Errorf("Wrong status got = %v, want %v", c.Writer.Status(), tt.args.status)
			}
			for _, err := range c.Errors {
				if err == nil {
					t.Errorf("Must contain error")
				}
				if !errors.Equal(errors.GetFirstNestedError(err), tt.args.err) {
					t.Errorf("JSONWithStatus() got = %v, want %v", err, tt.args.err)
				}
			}
		})
	}
}
