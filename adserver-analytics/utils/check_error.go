package utils

import "log"

// CheckError func
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
