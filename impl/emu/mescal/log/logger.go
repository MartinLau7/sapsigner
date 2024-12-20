package log

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Llongfile)
)

func Logger() *log.Logger {
	return logger
}
