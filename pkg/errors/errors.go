package errors

import (
	"fmt"
	"net/http"

	"git.daimler.com/cns/order-api/pkg/logger"
	"github.com/rs/zerolog"
)

const (
	Unexpected          = Kind(0)
	NoContent           = Kind(http.StatusNoContent)
	BadRequest          = Kind(http.StatusBadRequest)
	Unauthorized        = Kind(http.StatusUnauthorized)
	Forbidden           = Kind(http.StatusForbidden)
	NotFound            = Kind(http.StatusNotFound)
	InternalServerError = Kind(http.StatusInternalServerError)
)

type ErrorOption func(*Error)

func KindUnexpected() ErrorOption {
	return func(e *Error) {
		e.Kind = Unexpected
	}
}

func KindNoContent() ErrorOption {
	return func(e *Error) {
		e.Kind = NoContent
	}
}

func KindBadRequest() ErrorOption {
	return func(e *Error) {
		e.Kind = BadRequest
	}
}

func KindUnauthorized() ErrorOption {
	return func(e *Error) {
		e.Kind = Unauthorized
	}
}

func KindForbidden() ErrorOption {
	return func(e *Error) {
		e.Kind = Forbidden
	}
}

func KindNotFound() ErrorOption {
	return func(e *Error) {
		e.Kind = NotFound
	}
}

func KindInternalServerError() ErrorOption {
	return func(e *Error) {
		e.Kind = InternalServerError
	}
}

func WithKind(k Kind) ErrorOption {
	return func(e *Error) {
		switch k {
		case Unexpected:
			fallthrough
		case NoContent:
			fallthrough
		case BadRequest:
			fallthrough
		case Unauthorized:
			fallthrough
		case Forbidden:
			fallthrough
		case NotFound:
			fallthrough
		case InternalServerError:
			e.Kind = k
		default:
			e.Kind = Unexpected
		}
	}
}

func WithOp(o Op) ErrorOption {
	return func(e *Error) {
		e.Op = o
	}
}

func WithSeverity(s zerolog.Level) ErrorOption {
	return func(e *Error) {
		e.Severity = s
	}
}

func WithError(err error) ErrorOption {
	return func(e *Error) {
		e.Err = err
	}
}

func WithLevel(l zerolog.Level) ErrorOption {
	return func(e *Error) {
		e.Severity = l
	}
}

type Op string
type Kind int

func (k Kind) String() string {
	switch k {
	case Unexpected:
		return "Unexpected Error"
	case NoContent:
		return "No Content"
	case BadRequest:
		return "Bad Request"
	case Unauthorized:
		return "Unauthorized"
	case Forbidden:
		return "Forbidden"
	case NotFound:
		return "Not Found"
	case InternalServerError:
		return "Internal server error"
	default:
		return "Unexpected Error"
	}
}

func (k Kind) Int() int {
	if k == Unexpected {
		return http.StatusInternalServerError
	}
	return int(k)
}

type Error struct {
	Op       Op    // operation
	Kind     Kind  // category of errors
	Err      error // the wrapped error
	Severity zerolog.Level
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func Build(ops ...ErrorOption) *Error {
	e := &Error{}
	for _, op := range ops {
		op(e)
	}
	return e
}

// Creates a list with all recursive operations
func Ops(e *Error) []Op {
	res := []Op{e.Op}

	subErr, ok := e.Err.(*Error)
	if !ok {
		return res
	}

	res = append(res, Ops(subErr)...)
	return res
}

// Logs the error by level
func LogError(l logger.Interface, err error) {
	ll, ok := l.(*logger.Logger)
	if !ok {
		l.Error(err)
		return
	}

	sysErr, ok := err.(*Error)
	if !ok {
		ll.Error(err)
		return
	}

	ll.LogSysErr(sysErr.Severity, fmt.Sprintf("%s: %s", string(sysErr.Op), sysErr.Error()))
}

func Equal(e1 *Error, e2 *Error) bool {
	if e1 == nil && e2 == nil {
		return true
	}
	if (e1 != nil && e2 == nil) || (e1 == nil && e2 != nil) {
		return false
	}
	return e1.Err.Error() == e2.Err.Error() && e1.Kind == e2.Kind && e1.Op == e2.Op && e1.Severity == e2.Severity
}
