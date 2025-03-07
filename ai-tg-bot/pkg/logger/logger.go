package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	Info  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Infof(format string, args ...interface{}) {
	Info.Output(2, fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...interface{}) {
	Error.Output(2, fmt.Sprintf(format, args...))
}

func Debugf(format string, args ...interface{}) {
	Debug.Output(2, fmt.Sprintf(format, args...))
}
