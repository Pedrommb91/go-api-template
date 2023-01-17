package logger

import (
	"reflect"
	"testing"

	"github.com/rs/zerolog"
)

func TestNew(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name string
		args args
		want *Logger
	}{
		{
			name: "Debug logger",
			args: args{level: "debug"},
			want: New("debug"),
		},
		{
			name: "Warn logger",
			args: args{level: "warn"},
			want: New("warn"),
		},
		{
			name: "Error logger",
			args: args{level: "error"},
			want: New("error"),
		},
		{
			name: "Info logger",
			args: args{level: "info"},
			want: New("info"),
		},
		{
			name: "Wrong logger",
			args: args{level: "info"},
			want: New("wrong"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	l := New("debug")
	type fields struct {
		logger *zerolog.Logger
	}
	type args struct {
		message interface{}
		args    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Normal message and no args",
			fields: fields{logger: l.logger},
			args: args{
				message: "normal message",
			},
		},
		{
			name:   "Message with args",
			fields: fields{logger: l.logger},
			args: args{
				message: "Message with %d args",
				args:    []interface{}{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				logger: tt.fields.logger,
			}
			l.Debug(tt.args.message, tt.args.args...)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	l := New("info")
	type fields struct {
		logger *zerolog.Logger
	}
	type args struct {
		message string
		args    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Normal message and no args",
			fields: fields{logger: l.logger},
			args: args{
				message: "normal message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				logger: tt.fields.logger,
			}
			l.Info(tt.args.message, tt.args.args...)
		})
	}
}

func TestLogger_Warn(t *testing.T) {
	l := New("warn")
	type fields struct {
		logger *zerolog.Logger
	}
	type args struct {
		message string
		args    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Normal message and no args",
			fields: fields{logger: l.logger},
			args: args{
				message: "normal message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				logger: tt.fields.logger,
			}
			l.Warn(tt.args.message, tt.args.args...)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	l := New("error")
	type fields struct {
		logger *zerolog.Logger
	}
	type args struct {
		message interface{}
		args    []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Normal message and no args",
			fields: fields{logger: l.logger},
			args: args{
				message: "normal message",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				logger: tt.fields.logger,
			}
			l.Error(tt.args.message, tt.args.args...)
		})
	}
}

func TestLogger_LogSysErr(t *testing.T) {
	l := New("debug")
	type fields struct {
		logger *zerolog.Logger
	}
	type args struct {
		level   zerolog.Level
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Debug message",
			fields: fields{logger: l.logger},
			args: args{
				message: "debug message",
				level:   zerolog.DebugLevel,
			},
		},
		{
			name:   "Info message",
			fields: fields{logger: l.logger},
			args: args{
				message: "info message",
				level:   zerolog.InfoLevel,
			},
		},
		{
			name:   "Warn message",
			fields: fields{logger: l.logger},
			args: args{
				message: "warn message",
				level:   zerolog.WarnLevel,
			},
		},
		{
			name:   "Error message",
			fields: fields{logger: l.logger},
			args: args{
				message: "error message",
				level:   zerolog.ErrorLevel,
			},
		},
		{
			name:   "Fatal message",
			fields: fields{logger: l.logger},
			args: args{
				message: "fatal message",
				level:   zerolog.FatalLevel,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logger{
				logger: tt.fields.logger,
			}
			l.LogSysErr(tt.args.level, tt.args.message)
		})
	}
}
