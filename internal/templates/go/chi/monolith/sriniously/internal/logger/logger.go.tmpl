package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func NewLogger() zerolog.Logger {
	var logLevel zerolog.Level

	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	var writer io.Writer

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	writer = consoleWriter

	logger := zerolog.New(writer).
		Level(logLevel).
		With().
		Timestamp().
		Logger()

	return logger
}
