package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/juleur/mini-adserver/adserver-providers/handlers"
	"github.com/juleur/mini-adserver/adserver-providers/utils"
)

func main() {
	log.Println("Adserver Providers")
	jsonFilename := flag.String("f", "campaigns.json", "")
	flag.Parse()

	jsonData := utils.JSONFileOpener(*jsonFilename)

	http.Handle("/ad", handlers.AdserverHandler(jsonData))
	http.ListenAndServe(":2323", nil)
}
