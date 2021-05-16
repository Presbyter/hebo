package log

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type log struct {
	output io.Writer
	m      sync.RWMutex
}

func (l *log) Error(v ...interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	level := ERROR
	write(l.output, fmt.Sprint(v...), level)
}

func (l *log) Errorf(format string, v ...interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	level := ERROR
	write(l.output, fmt.Sprintf(format, v...), level)
}

func (l *log) Debug(v ...interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	level := DEBUG
	write(l.output, fmt.Sprint(v...), level)
}

func (l *log) Debugf(format string, v ...interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	level := DEBUG
	write(l.output, fmt.Sprintf(format, v...), level)
}

func (l *log) Info(v ...interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	level := INFO
	write(l.output, fmt.Sprint(v...), level)
}

func (l *log) Infof(format string, v ...interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	level := INFO
	write(l.output, fmt.Sprintf(format, v...), level)
}

func (l *log) Warn(v ...interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	level := WARN
	write(l.output, fmt.Sprint(v...), level)
}

func (l *log) Warnf(format string, v ...interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	level := WARN
	write(l.output, fmt.Sprintf(format, v...), level)
}

func (l *log) Fatal(v ...interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	level := FATAL
	write(l.output, fmt.Sprint(v...), level)
}

func (l *log) Fatalf(format string, v ...interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	level := FATAL
	write(l.output, fmt.Sprintf(format, v...), level)
}

func (l *log) SetOutput(path string) error {
	var file *os.File
	var err error
	if _, err = os.Stat(path); os.IsNotExist(err) {
		// 创建文件
		if file, err = os.Create(path); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if file, err = os.Open(path); err != nil {
		return err
	}
	l.output = file
	return nil
}

func write(wr io.Writer, str string, level LogLevel) {
	w := bufio.NewWriter(wr)
	w.WriteString(level.String())
	w.WriteString(str)
	w.WriteByte('\n')
}

func New() *log {
	return Default
}

var (
	Default = &log{
		output: os.Stdout,
		m:      sync.RWMutex{},
	}
)
