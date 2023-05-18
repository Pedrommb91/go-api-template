package errors

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
)

var NewUUID = uuid.NewV4

const (
	Unexpected          = Kind(0)
	NoContent           = Kind(http.StatusNoContent)
	BadRequest          = Kind(http.StatusBadRequest)
	Unauthorized        = Kind(http.StatusUnauthorized)
	Forbidden           = Kind(http.StatusForbidden)
	NotFound            = Kind(http.StatusNotFound)
	RequestTimeout      = Kind(http.StatusRequestTimeout)
	Conflict            = Kind(http.StatusConflict)
	InternalServerError = Kind(http.StatusInternalServerError)
	BadGateway          = Kind(http.StatusBadGateway)
)

var kindLookUp = map[Kind]string{
	Unexpected:          "Unexpected Error",
	NoContent:           "No Content",
	BadRequest:          "Bad Request",
	BadGateway:          "Bad Gateway",
	Unauthorized:        "Unauthorized",
	Forbidden:           "Forbidden",
	NotFound:            "Not Found",
	RequestTimeout:      "Request Timeout",
	Conflict:            "Conflict",
	InternalServerError: "Internal server error",
}

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

func KindBadGateway() ErrorOption {
	return func(e *Error) {
		e.Kind = BadGateway
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

func KindConflict() ErrorOption {
	return func(e *Error) {
		e.Kind = Conflict
	}
}

func KindRequestTimout() ErrorOption {
	return func(e *Error) {
		e.Kind = RequestTimeout
	}
}

func KindInternalServerError() ErrorOption {
	return func(e *Error) {
		e.Kind = InternalServerError
	}
}

func WithKind(k Kind) ErrorOption {
	return func(e *Error) {
		_, ok := kindLookUp[k]
		if ok {
			e.Kind = k
		} else {
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
		er, ok := err.(*Error)
		if ok {
			e.ID = er.ID
		} else {
			e.ID = generateID()
		}

		e.Err = err
	}
}

// WithNestedErrorCopy copy id, kind, severity and message from the input error
func WithNestedErrorCopy(err error) ErrorOption {
	return func(e *Error) {
		er, ok := err.(*Error)
		if ok {
			e.ID = er.ID
			e.Kind = er.Kind
			e.Severity = er.Severity
			e.Message = er.Message
		} else {
			e.ID = generateID()
		}

		e.Err = err
	}
}

func WithID(id string) ErrorOption {
	return func(e *Error) {
		e.ID = id
	}
}

func WithMessage(msg string) ErrorOption {
	return func(e *Error) {
		e.Message = msg
	}
}

type Op string
type Kind int

func (k Kind) String() string {
	s, ok := kindLookUp[k]
	if ok {
		return s
	}

	return "Unexpected Error"
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
	Message  string
	Severity zerolog.Level
	ID       string
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func Build(ops ...ErrorOption) *Error {
	e := newCustomErrorWithDefaults()

	for _, op := range ops {
		op(e)
	}

	LogErrorWithSeverity(logger.New(e.Severity.String()), e.Severity, e)

	return e
}

func newCustomErrorWithDefaults() *Error {
	return &Error{
		Op:       "No operation found",
		Err:      fmt.Errorf("no error"),
		Message:  "No message",
		Kind:     Unexpected,
		Severity: zerolog.WarnLevel,
		ID:       generateID(),
	}
}

// Creates a list with all recursive operations.
func Ops(e *Error) []Op {
	res := []Op{e.Op}

	subErr, ok := e.Err.(*Error)
	if !ok {
		return res
	}

	return append(res, Ops(subErr)...)
}

func GetFirstNestedError(e error) error {
	err, ok := e.(*Error)
	if !ok {
		return e
	}

	for {
		subErr, ok := err.Err.(*Error)
		if ok {
			err = subErr
		} else {
			break
		}
	}

	return err
}

// Logs the error by level.
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

	ll.LogSysErr(sysErr.Severity, fmt.Sprintf("[%s]%s: %s", sysErr.ID, string(sysErr.Op), sysErr.Error()))
}

func LogErrorWithSeverity(l logger.Interface, severity zerolog.Level, err error) {
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

	ll.LogSysErr(severity, fmt.Sprintf("[%s]%s: %s - %s", sysErr.ID, string(sysErr.Op), sysErr.Message, sysErr.Err.Error()))
}

func Equal(e1 error, e2 error) bool {
	if e1 == nil && e2 == nil {
		return true
	}

	if (e1 != nil && e2 == nil) || (e1 == nil && e2 != nil) {
		return false
	}

	err1, ok1 := e1.(*Error)
	err2, ok2 := e2.(*Error)

	if ok1 != ok2 {
		return false
	}
	if !ok1 && !ok2 {
		return reflect.DeepEqual(e1, e2)
	}

	if (err1 == nil && err2 == nil) || (e1 == nil && err2 == nil) || (err1 == nil && e2 == nil) {
		return true
	}

	if (err1 != nil && err2 == nil) || (err1 == nil && err2 != nil) {
		return false
	}

	return err1.Err.Error() == err2.Err.Error() &&
		err1.Kind == err2.Kind &&
		err1.Op == err2.Op &&
		err1.Severity == err2.Severity &&
		err1.Message == err2.Message &&
		err1.ID == err2.ID
}

func generateID() string {
	idstr := NewUUID().String()
	return idstr
}
