package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/juleur/mini-adserver/adserver-providers/utils"

	"github.com/juleur/mini-adserver/adserver-analytics/models"
)

// GetSumPlacementHandler s
func GetSumPlacementHandler(cm *models.CampaignManager) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			placementID := r.URL.Query().Get("placement")
			res := cm.StoredData[placementID]
			cs, err := json.Marshal(res)
			utils.CheckError(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(cs)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
}
