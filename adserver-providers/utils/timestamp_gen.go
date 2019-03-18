package utils

import (
	"fmt"
	"time"
)

// TimestampGenerator generates timestamp
func TimestampGenerator() string {
	t := time.Now()
	return fmt.Sprintf("%02d-%02d-%d %02d:%02d:%02d",
		t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
}
