package utils

import (
	"log"
	"net/http"
)

// CheckError func
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// HttpError func
func HTTPError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
