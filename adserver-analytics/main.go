package main

import (
	"log"
	"net/http"

	"github.com/juleur/mini-adserver/adserver-analytics/handlers"
	"github.com/juleur/mini-adserver/adserver-analytics/utils"
)

func main() {
	log.Println("Adserver Analytics")
	campaignManager := utils.NewCampaignManager()
	utils.StoreDataAtBoot(campaignManager)

	http.Handle("/campaign_provided", handlers.CampaignDoneHandler(campaignManager))
	http.Handle("/sum", handlers.GetSumCampaignsHandler(campaignManager))
	http.Handle("/sum_placement", handlers.GetSumPlacementHandler(campaignManager))
	http.ListenAndServe(":3030", nil)
}
