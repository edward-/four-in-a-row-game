package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger interface{
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
}

type logger struct{}

func NewLogger() Logger {
	logger := new(logger)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return logger
}

func (l *logger) Debug(msg string) {
	log.Debug().Msg(msg)
}

func (l *logger) Info(msg string) {
	log.Info().Msg(msg)
}

func (l *logger) Warn(msg string) {
	log.Warn().Msg(msg)
}

func (l *logger) Error(msg string) {
	log.Error().Msg(msg)
}