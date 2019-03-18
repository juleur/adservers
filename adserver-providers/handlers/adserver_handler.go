package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/juleur/mini-adserver/adserver-providers/models"
	"github.com/juleur/mini-adserver/adserver-providers/utils"
)

func redirectToAnalytics(c models.Campaign) error {
	campaign := models.CampaignAnalytic{
		Timestamp:   utils.TimestampGenerator(),
		PlacementID: c.Placements[0],
		CampaignID:  c.ID,
		Price:       c.Price,
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(campaign)
	//handles errors
	_, err := http.Post("http://127.0.0.1:3030/campaign_provided", "application/json", buf)
	if err != nil {
		return err
	}
	return nil
}

// parsePostRequest decodes body request and gets url query
// then make it into BodyRequest struct
func parsePostRequest(r *http.Request) (models.BodyRequest, error) {
	bodyReq := models.BodyRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return models.BodyRequest{}, err
	}
	err = json.Unmarshal(body, &bodyReq)
	if err != nil {
		return models.BodyRequest{}, err
	}
	bodyReq.PlacementID = r.URL.Query().Get("placement")
	return bodyReq, nil
}

// AdserverHandler is the http POST response
func AdserverHandler(jsonData []models.Campaign) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			reqParsed, err := parsePostRequest(r)
			utils.HTTPError(w, err)
			campaign := utils.GetCampaign(reqParsed, jsonData)
			serResponse := &models.Campaign{
				ID:      campaign.ID,
				Content: campaign.Content,
			}
			cpgn, err := json.Marshal(serResponse)
			utils.HTTPError(w, err)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(cpgn)
			err = redirectToAnalytics(campaign)
			if err != nil {
				log.Println(err)
			}
		default:
			w.WriteHeader(http.StatusForbidden)
		}
	})
}
