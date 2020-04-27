package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Stamp})
}

func Info(msg string) {
	log.Info().Caller(1).Msg(msg)
}

func Infof(format string, v ...interface{}) {
	log.Info().Caller(1).Msgf(format, v...)
}

func Debug(msg string) {
	log.Debug().Caller(1).Msg(msg)
}

func Debugf(format string, v ...interface{}) {
	log.Debug().Caller(1).Msgf(format, v...)
}

func Error(err error) {
	log.Error().Caller(1).Msgf("%+v", err)
}

func Fatal(err error) {
	log.Fatal().Caller(1).Msgf("%+v", err)
}
