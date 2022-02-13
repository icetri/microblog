package pkg

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var (
	logger    = zerolog.New(os.Stdout)
	loggerErr = zerolog.New(os.Stderr)
)

func LogError(err error) {
	loggerErr.Err(err).Time("time", time.Now()).Send()
}

func FatalError(err error) {
	loggerErr.Fatal().Err(err).Time("time", time.Now()).Send()
}

func LogInfo(msg string) {
	logger.Info().Time("time", time.Now()).Msg(msg)
}
