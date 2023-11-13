package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

var fw *fileWriter

func Init() bool {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// set output
	var err error
	prefix := filepath.Base(os.Args[0])

	fw, err = newFileWriter("log/", prefix)
	if err != nil {
		panic(err)
	}

	log.Logger = log.Output(fw).With().Caller().Logger()
	gin.DefaultWriter = fw

	return true
}

func Close() error {
	return fw.Close()
}

func Info() *zerolog.Event {
	return log.Info()
}

func Error() *zerolog.Event {
	return log.Error()
}
