package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Interface -.
type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// Logger -.
type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

// New -.
func New(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	//zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger().Level(l)

	return &Logger{
		logger: &logger,
	}
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.msg("debug", message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.log("info", message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.log("warn", message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	l.msg("error", message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg("fatal", message, args...)
	os.Exit(1)
}

func (l *Logger) log(level string, message string, args ...interface{}) {
	var m string
	if len(args) == 0 {
		m = message
	} else {
		m = fmt.Sprintf(message, args...)
	}

	switch strings.ToLower(level) {
	case "error":
		l.logger.Error().Msg(m)
	case "warn":
		l.logger.Warn().Msg(m)
	case "info":
		l.logger.Info().Msg(m)
	case "debug":
		l.logger.Debug().Msg(m)
	case "fatal":
		l.logger.Fatal().Msg(m)
	default:
		l.log(l.logger.GetLevel().String(), message, args...)
	}
}

func (l *Logger) msg(level string, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(level, msg.Error(), args...)
	case string:
		l.log(level, msg, args...)
	default:
		l.log(l.logger.GetLevel().String(), fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}

func (l *Logger) LogSysErr(level zerolog.Level, message string) {
	switch level {
	case zerolog.ErrorLevel:
		l.Error(message)
	case zerolog.WarnLevel:
		l.Warn(message)
	case zerolog.InfoLevel:
		l.Info(message)
	case zerolog.DebugLevel:
		l.Debug(message)
	default:
		l.Error(message)
	}
}
