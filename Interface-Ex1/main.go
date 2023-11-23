package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

const format = "%v : INFO : %s"

type Logger interface {
	Info(string)
}

type ScreenLogger struct{}

type FileLogger struct {
	File *os.File
}

func (ScreenLogger) Info(message string) {
	fmt.Printf(format, time.Now(), message)
}

func (l FileLogger) Info(message string) {
	fmt.Fprintf(l.File, format, time.Now(), message)
}

func NewFileLogger(filename string) *FileLogger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	return &FileLogger{file}
}

func main() {
	var log Logger
	//log = &ScreenLogger{}
	log = NewFileLogger("log.txt")
	log.Info("Hello World\n")
	fmt.Println(reflect.TypeOf(log))
}
