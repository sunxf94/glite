package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("hello world")
	log.Log()
}

func TestFile(t *testing.T) {
	f, err := os.OpenFile("./tmp.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		t.Error(err)
		return
	}

	wg := sync.WaitGroup{}
	for i := 10000; i < 20000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			msg := fmt.Sprintf("id:\t%d, time: %d\n", id, time.Now().UnixNano())
			if _, err = f.Write([]byte(msg)); err != nil {
				t.Error(err)
			}
		}(i)
	}

	wg.Wait()
}

func TestAny(t *testing.T) {
	t.Log(os.Args)
	t.Log(os.Args[0])
	t.Log(filepath.Base(os.Args[0]))
	hostname, err := os.Hostname()
	t.Log(hostname, err)
}
