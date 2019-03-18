package utils

import (
	"log"
	"os"
)

// TextOpener s
func TextOpener() *os.File {
	file, err := os.OpenFile("analytics.txt", os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
