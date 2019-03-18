package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/juleur/mini-adserver/adserver-analytics/utils"

	"github.com/juleur/mini-adserver/adserver-analytics/models"
)

// CampaignDoneHandler s
func CampaignDoneHandler(cm *models.CampaignManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			body, err := ioutil.ReadAll(r.Body)
			utils.CheckError(err)
			c := models.CampaignAnalytic{}
			json.Unmarshal(body, &c)
			// handler err
			fmt.Printf("%s,%s,%s,%.2f\n", c.Timestamp, c.PlacementID, c.CampaignID, c.Price)
			cm.File.Write(body)
			cm.SumTotal += c.Price
			utils.StoreCampaigns(c, cm.StoredData)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
}
