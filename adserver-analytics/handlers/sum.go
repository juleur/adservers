package handlers

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/juleur/mini-adserver/adserver-analytics/models"
	"github.com/juleur/mini-adserver/adserver-analytics/utils"
)

// GetSumCampaignsHandler gets how much adserver made profit
func GetSumCampaignsHandler(cm *models.CampaignManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fround := math.Round(cm.SumTotal*100) / 100
			sum, err := json.Marshal(fround)
			utils.CheckError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(sum)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
}
