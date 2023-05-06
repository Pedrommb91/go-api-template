package errors

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Pedrommb91/go-api-template/pkg/logger"
	"github.com/rs/zerolog"
)

func TestBuild(t *testing.T) {
	id := "dummy-id"
	type args struct {
		ops []ErrorOption
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "When dont send an error",
			args: args{[]ErrorOption{WithID(id)}},
			want: &Error{
				Kind:     Unexpected,
				Err:      fmt.Errorf("no error"),
				Severity: zerolog.DebugLevel,
				ID:       id,
				Op:       "No operation found",
				Message:  "No message",
			},
		},
		{
			name: "When send an error",
			args: args{[]ErrorOption{
				KindForbidden(),
				WithSeverity(zerolog.WarnLevel),
			}},
			want: &Error{
				Err:      fmt.Errorf("no error"),
				Kind:     Forbidden,
				ID:       id,
				Severity: zerolog.WarnLevel,
				Op:       "No operation found",
				Message:  "No message",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Build(tt.args.ops...); !Equal(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		Op       Op
		Kind     Kind
		Err      error
		Severity zerolog.Level
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "When Internal Server Error",
			fields: fields{Err: &Error{
				Err: fmt.Errorf("Internal Server Error"),
			}},
			want: string("Internal Server Error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Op:       tt.fields.Op,
				Kind:     tt.fields.Kind,
				Err:      tt.fields.Err,
				Severity: tt.fields.Severity,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOps(t *testing.T) {
	err := Build(
		WithOp(Op("TestOps")),
		WithError(fmt.Errorf("Internal Server Error")),
		WithSeverity(zerolog.ErrorLevel),
	)

	err1 := Build(
		WithOp(Op("Nested.Error")),
		WithError(err),
		WithSeverity(zerolog.ErrorLevel),
	)

	type args struct {
		e *Error
	}
	tests := []struct {
		name string
		args args
		want []Op
	}{
		{
			name: "When dont have Ops",
			args: args{err},
			want: []Op{err.Op},
		},
		{
			name: "When have nested Error",
			args: args{err1},
			want: []Op{err1.Op, err.Op},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ops(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ops() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogError(t *testing.T) {
	err := Build(
		WithOp(Op(fmt.Errorf("").Error())),
		WithError(fmt.Errorf("Internal Server Error")),
		WithSeverity(zerolog.ErrorLevel),
	)
	type args struct {
		l   logger.Interface
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "When LogError",
			args: args{logger.New("debug"), err},
		},
		{
			name: "When Wrong Error",
			args: args{logger.New("debug"), fmt.Errorf("wrong error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LogError(tt.args.l, tt.args.err)
		})
	}
}

func TestEqual(t *testing.T) {
	err := Build(
		WithOp(Op(fmt.Errorf("").Error())),
		WithError(fmt.Errorf("Internal Server Error")),
		WithSeverity(zerolog.ErrorLevel),
	)

	err2 := Build(
		WithOp(Op(fmt.Errorf("").Error())),
		WithError(fmt.Errorf("BadRequest")),
		KindBadRequest(),
	)
	type args struct {
		e1 *Error
		e2 *Error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "When error are nil",
			args: args{e1: nil, e2: nil},
			want: true,
		},
		{
			name: "When error are different",
			args: args{nil, err2},
			want: false,
		},
		{
			name: "When error are different",
			args: args{err, err2},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.e1, tt.args.e2); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKind_String(t *testing.T) {
	tests := []struct {
		name string
		k    Kind
		want string
	}{
		{
			name: "Unexpected Error",
			k:    Unexpected,
			want: "Unexpected Error",
		},
		{
			name: "No content error",
			k:    NoContent,
			want: "No Content",
		},
		{
			name: "Bad request error",
			k:    BadRequest,
			want: "Bad Request",
		},
		{
			name: "Bad gateway error",
			k:    BadGateway,
			want: "Bad Gateway",
		},
		{
			name: "Unauthorized error",
			k:    Unauthorized,
			want: "Unauthorized",
		},
		{
			name: "Forbidden error",
			k:    Forbidden,
			want: "Forbidden",
		},
		{
			name: "not found error",
			k:    NotFound,
			want: "Not Found",
		},
		{
			name: "time out error",
			k:    RequestTimeout,
			want: "Request Timeout",
		},
		{
			name: "Internal server error",
			k:    InternalServerError,
			want: "Internal server error",
		},
		{
			name: "An Unexpected Error",
			k:    800,
			want: "Unexpected Error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.String(); got != tt.want {
				t.Errorf("Kind.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKind_Int(t *testing.T) {
	tests := []struct {
		name string
		k    Kind
		want int
	}{
		{
			name: "Unexpected ERROR",
			k:    Unexpected,
			want: 500,
		},
		{
			name: "Other type of error",
			k:    BadRequest,
			want: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.k.Int(); got != tt.want {
				t.Errorf("Kind.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithKind(t *testing.T) {
	id := "dummy-id"
	type args struct {
		k Kind
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "When unexpected",
			args: args{
				k: Unexpected,
			},
			want: Build(WithID(id)),
		},
		{
			name: "When Not Found",
			args: args{
				k: NotFound,
			},
			want: Build(KindNotFound(), WithID(id)),
		},

		{
			name: "When Internal Server Error",
			args: args{
				k: InternalServerError,
			},
			want: Build(KindInternalServerError(), WithID(id)),
		},
		{
			name: "When Bad request",
			args: args{
				k: BadRequest,
			},
			want: Build(KindBadRequest(), WithID(id)),
		},
		{
			name: "When Bad gateway",
			args: args{
				k: BadGateway,
			},
			want: Build(KindBadGateway(), WithID(id)),
		},
		{
			name: "When Unauthorized",
			args: args{
				k: Unauthorized,
			},
			want: Build(KindUnauthorized(), WithID(id)),
		},
		{
			name: "When No Content",
			args: args{
				k: NoContent,
			},
			want: Build(KindNoContent(), WithID(id)),
		},
		{
			name: "When Unexistent kind",
			args: args{
				k: Kind(-1),
			},
			want: Build(WithID(id)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Build(WithKind(tt.args.k), WithID(id)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFirstNestedError(t *testing.T) {
	id := "dummy-id"
	firstErr := fmt.Errorf("first error")
	type args struct {
		e error
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "nested error return first error",
			args: args{
				e: Build(
					WithError(Build(WithError(firstErr), WithID(id))),
				),
			},
			wantErr: Build(WithError(firstErr), WithID(id)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetFirstNestedError(tt.args.e); !Equal(err, tt.wantErr) {
				t.Errorf("GetFirstNestedError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
