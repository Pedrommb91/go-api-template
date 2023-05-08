package errors

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Pedrommb91/go-api-template/pkg/clock/mocks"
	"github.com/Pedrommb91/go-api-template/pkg/errors"
	"github.com/gin-gonic/gin"
)

func TestAbortWithStatusJSON(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Valid costum error, context is aborted and error returned",
			args: args{
				err: errors.Build(),
			},
		},
		{
			name: "Valid native error, context is aborted and error returned",
			args: args{
				err: errors.Build(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clockMock := mocks.NewClock(t)
			clockMock.On("Now").Return(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)).Maybe()
			c, _ := gin.CreateTestContext(httptest.NewRecorder())

			e := NewErrorSender(clockMock)
			e.AbortWithStatusJSON(c, tt.args.err)
			if !c.IsAborted() {
				t.Errorf("Context not aborted")
			}
			for _, err := range c.Errors {
				if err == nil || err.Err == nil {
					t.Errorf("Must contain error")
				}
				if !errors.Equal(errors.GetFirstNestedError(err.Err), tt.args.err) {
					t.Errorf("AbortWithStatusJSON() got = %v, want %v", err, tt.args.err)
				}
			}
		})
	}
}

func TestJSON(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Valid costum error, context is not aborted and error returned",
			args: args{
				err: errors.Build(),
			},
		},
		{
			name: "Valid native error, context is not aborted and error returned",
			args: args{
				err: fmt.Errorf("error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clockMock := mocks.NewClock(t)
			clockMock.On("Now").Return(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)).Maybe()

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			e := NewErrorSender(clockMock)
			e.JSON(c, tt.args.err)

			if c.IsAborted() {
				t.Errorf("Context aborted")
			}
			for _, err := range c.Errors {
				if err == nil || err.Err == nil {
					t.Errorf("Must contain error")
				}
				if !errors.Equal(errors.GetFirstNestedError(err.Err), tt.args.err) {
					t.Errorf("JSON() got = %v, want %v", err, tt.args.err)
				}
			}
		})
	}
}

func TestJSONWithStatus(t *testing.T) {
	type args struct {
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
				err:    errors.Build(),
				status: 400,
			},
		},
		{
			name: "Valid native error, context is not aborted and error returned",
			args: args{
				err:    fmt.Errorf("error"),
				status: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clockMock := mocks.NewClock(t)
			clockMock.On("Now").Return(time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)).Maybe()

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			e := NewErrorSender(clockMock)
			e.JSONWithStatus(c, tt.args.err, tt.args.status)
			if c.IsAborted() {
				t.Errorf("Context aborted")
			}
			if c.Writer.Status() != tt.args.status {
				t.Errorf("Wrong status got = %v, want %v", c.Writer.Status(), tt.args.status)
			}
			for _, err := range c.Errors {
				if err == nil || err.Err == nil {
					t.Errorf("Must contain error")
				}
				if !errors.Equal(errors.GetFirstNestedError(err.Err), tt.args.err) {
					t.Errorf("JSONWithStatus() got = %v, want %v", err, tt.args.err)
				}
			}
		})
	}
}
