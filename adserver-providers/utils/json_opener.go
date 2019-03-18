package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/juleur/mini-adserver/adserver-providers/models"
)

// campaignArray converts map into []Campaign structs
func campaignArray(campaigns models.Campaigns) []models.Campaign {
	cs := []models.Campaign{}
	for pID, campaign := range campaigns.Campaigns {
		campaign.ID = pID
		cs = append(cs, campaign)
	}
	return cs
}

// JSONFileOpener func
func JSONFileOpener(filename string) []models.Campaign {
	campaigns := models.Campaigns{}
	content, err := ioutil.ReadFile(filename)
	CheckError(err)
	err = json.Unmarshal(content, &campaigns)
	CheckError(err)
	cmpgs := campaignArray(campaigns)
	return cmpgs
}
