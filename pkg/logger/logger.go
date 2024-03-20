package logger

import (
	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(error, string)
	Fatal(error, string, string)
}

type logger struct {
	log zerolog.Logger
}

func NewLogger() Logger {
	l := setup()
	logger := new(logger)
	logger.log = l
	return logger
}

func (l *logger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l *logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *logger) Warn(msg string) {
	l.log.Warn().Msg(msg)
}

func (l *logger) Error(err error, msg string) {
	l.log.Err(err).Msg(msg)
}

func (l *logger) Fatal(err error, service string, msg string) {
	l.log.Fatal().
		Err(err).
		Str("service", service).
		Msgf("Fatal error: %s", service)
}
