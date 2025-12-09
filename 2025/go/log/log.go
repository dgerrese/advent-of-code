package log

import (
	"log"
	"os"
)

func CreateLogFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
}

var DefaultFile, _ = CreateLogFile("AOC2025.log")

func NewLogger(prefix string) *log.Logger {
	return log.New(DefaultFile, prefix, log.Ltime)
}
