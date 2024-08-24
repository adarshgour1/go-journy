package utils

import (
	"log"
	"os"
)

func NewLogger(filename string) *log.Logger {
	buf, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	lg := log.New(buf, "", log.Lmicroseconds|log.Lshortfile)
	return lg

}
