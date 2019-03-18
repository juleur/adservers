package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/juleur/mini-adserver/adserver-providers/handlers"
	"github.com/juleur/mini-adserver/adserver-providers/utils"
)

func main() {
	log.Println("Adserver Providers Started")
	jsonFilename := flag.String("f", "campaigns.json", "")
	flag.Parse()

	jsonData := utils.JSONFileOpener(*jsonFilename)
	analyticServIPAddr := "http://192.168.0.122"

	http.Handle("/ad", handlers.AdserverHandler(jsonData, analyticServIPAddr))
	http.ListenAndServe(":2323", nil)
	log.Println("Adserver Providers Stopped")
}
