package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/juleur/mini-adserver/adserver-providers/models"
	"github.com/juleur/mini-adserver/adserver-providers/utils"
)

func redirectToAnalytics(c models.Campaign, analyticServIPAddr string) error {
	campaign := models.CampaignAnalytic{
		Timestamp:   utils.TimestampGenerator(),
		PlacementID: c.Placements[0],
		CampaignID:  c.ID,
		Price:       c.Price,
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(campaign)
	//handles errors
	if err != nil {
		return err
	}
	_, err = http.Post(fmt.Sprintf("%s:3030/campaign_provided", analyticServIPAddr), "application/json", buf)
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
func AdserverHandler(jsonData []models.Campaign, analyticServIPAddr string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			reqParsed, err := parsePostRequest(r)
			utils.HTTPError(w, err)
			campaign, notFound := utils.GetCampaign(reqParsed, jsonData)
			if notFound {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			serResponse := &models.Campaign{
				ID:      campaign.ID,
				Content: campaign.Content,
			}
			cpgn, err := json.Marshal(serResponse)
			utils.HTTPError(w, err)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(cpgn)
			err = redirectToAnalytics(campaign, analyticServIPAddr)
			if err != nil {
				log.Println(err)
			}
		default:
			w.WriteHeader(http.StatusForbidden)
		}
	})
}
