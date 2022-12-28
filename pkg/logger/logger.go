package logger

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var (
	prefix     string
	filePath   = ""
	fileWriter logFileWriter
	nameFormat = "2006_01_02"
)

func SetPath(path string) {
	filePath = path
}

func Init() bool {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var w io.Writer
	if filePath == "" {
		w = zerolog.ConsoleWriter{Out: os.Stdout}
	} else {
		// set output
		prefix = filepath.Base(os.Args[0])
		hostName, err := os.Hostname()
		if err == nil {
			prefix = prefix + "_" + hostName
		}
		f := fileWriter.createFile(prefix, time.Now())
		if f == nil {
			return false
		}
		gin.DefaultWriter = f
		fileWriter.f = f
		fileWriter.stdout = false

		w = &fileWriter
	}
	log.Logger = log.Output(w).With().Caller().Logger()

	return true
}

func Info() *zerolog.Event {
	return log.Info()
}

func Error() *zerolog.Event {
	return log.Error()
}

type logFileWriter struct {
	mut    sync.Mutex
	f      *os.File
	stdout bool
}

func (l *logFileWriter) createFile(prefix string, now time.Time) *os.File {
	timeStr := now.Format(nameFormat)
	name := fmt.Sprintf("%s%s_%s.log", filePath, prefix, timeStr)
	file, _ := os.OpenFile(name, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	return file
}

func (l *logFileWriter) Write(p []byte) (n int, err error) {
	// 控制台输出
	if l.stdout {
		os.Stderr.Write(p)
	}
	// 文件输出
	tf := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&l.f)))
	f := (*os.File)(tf)
	if f != nil {
		l.mut.Lock()
		n, err := f.Write(p)
		l.mut.Unlock()
		return n, err
	}
	return 0, errors.New("logwritter error")
}

func (l *logFileWriter) Close() {
	tf := atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&l.f)), nil)
	f := (*os.File)(tf)
	if f != nil {
		f.Sync()
		f.Close()
	}
}
