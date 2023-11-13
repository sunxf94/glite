package logger

import (
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	errFileEmptyPointer = errors.New("empty pointer of file")
)

type fileWriter struct {
	path     string
	prefix   string
	file     *os.File
	filename string
	stdout   bool

	quit chan struct{}
}

func newFileWriter(path, prefix string) (*fileWriter, error) {
	fw := &fileWriter{
		path:   path,
		prefix: prefix,
	}
	_, err := fw.newFile()

	go fw.monitor()

	return fw, err
}

func (w *fileWriter) Write(p []byte) (int, error) {
	// 输出到控制台 os.Stderr相对于os.Stdout，可以直接输出而不是等待换行符
	if w.stdout {
		_, _ = os.Stderr.Write(p) // 写入控制台时 还是要写入到日志文件中
	}

	f := w.file // f判空与f.Write() 不是原子操作
	if f == nil {
		return 0, errFileEmptyPointer
	}

	return f.Write(p)
}

func (w *fileWriter) Close() error {
	f := w.file // f判空与f.Write() 不是原子操作
	if f == nil {
		return errFileEmptyPointer
	}
	defer func() { w.quit <- struct{}{} }()

	if err := f.Sync(); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func (w *fileWriter) newFile() (*os.File, error) {
	w.filename = w.newFilename()
	f, err := os.OpenFile(w.filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err == nil {
		oldFile := w.file
		w.file = f
		_ = oldFile.Close() // 根据go底层源码 这个操作应该是并发安全的
	}

	return f, err
}

func (w *fileWriter) newFilename() string {
	timeStr := time.Now().Format("2006-01-02")
	hostname, _ := os.Hostname()
	return fmt.Sprintf("%s%s_%s_%s.log", w.path, w.prefix, hostname, timeStr)
}

func (w *fileWriter) monitor() {
	for {
		select {
		case <-time.Tick(time.Second):
			if !w.needChanged() {
				continue
			}
			if _, err := w.newFile(); err != nil {
				panic(err) // 创建日志文件失败时 服务应该有反应
			}
		case <-w.quit:
			break
		}
	}
}

func (w *fileWriter) needChanged() bool {
	// 文件名中的日期更新时
	if newFilename := w.newFilename(); w.filename != newFilename {
		return true
	}
	// 文件名不存在时
	if _, err := os.Stat(w.filename); err != nil && !os.IsExist(err) {
		return true
	}

	return false
}
