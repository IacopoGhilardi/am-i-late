package logger

import (
	"fmt"
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	log  zerolog.Logger
	once sync.Once
)

func initLogger() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	log = zerolog.New(consoleWriter).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func Info(msg string, args ...interface{})  { logWithErr(zerolog.InfoLevel, msg, args...) }
func Debug(msg string, args ...interface{}) { logWithErr(zerolog.DebugLevel, msg, args...) }
func Warn(msg string, args ...interface{})  { logWithErr(zerolog.WarnLevel, msg, args...) }
func Error(msg string, args ...interface{}) { logWithErr(zerolog.ErrorLevel, msg, args...) }
func Fatal(msg string, args ...interface{}) { logWithErr(zerolog.FatalLevel, msg, args...); os.Exit(1) }

func logWithErr(level zerolog.Level, msg string, args ...interface{}) {
	once.Do(initLogger)

	l := log.WithLevel(level)
	formattedMsg := fmt.Sprintf(msg, args...)

	// cerchiamo il primo argomento che sia un errore
	for _, arg := range args {
		if err, ok := arg.(error); ok {
			l = l.Err(err)
			break
		}
	}

	l.Msg(formattedMsg)
}
