package logger

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/edward-/four-in-a-row-game/build"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once

var log zerolog.Logger

func setup() zerolog.Logger {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}

		log = zerolog.New(output).
			With().
			Timestamp().
			Str("tag", build.Version).
			Str("hash_commit", build.HashCommit).
			Logger()
	})

	return log
}
