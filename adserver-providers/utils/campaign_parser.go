package utils

import (
	"sort"

	"github.com/juleur/mini-adserver/adserver-providers/models"
)

func isDeviceOk(bodyDevice string, campaignDevices []string) bool {
	for _, v := range campaignDevices {
		if v == bodyDevice {
			return true
		}
	}
	return false
}

func isCountryOk(bodyCountry string, campaignCountries []string) bool {
	for _, v := range campaignCountries {
		if v == bodyCountry {
			return true
		}
	}
	return false
}

func isPlacementOk(bodyPlacementID string, campaignPlacements []string) bool {
	for _, v := range campaignPlacements {
		if v == bodyPlacementID {
			return true
		}
	}
	return false
}

// GetCampaign s
func GetCampaign(bodyReq models.BodyRequest, jsonData []models.Campaign) (models.Campaign, bool) {
	cs := []models.Campaign{}
	bestCampaign := models.Campaign{}
	for _, campaignData := range jsonData {
		pOK := isPlacementOk(bodyReq.PlacementID, campaignData.Placements)
		cOK := isCountryOk(bodyReq.Country, campaignData.Countries)
		dOK := isDeviceOk(bodyReq.Device, campaignData.Devices)
		if pOK && cOK && dOK {
			campaignData.Placements[0] = bodyReq.PlacementID
			cs = append(cs, campaignData)
		}
	}
	if len(cs) == 0 {
		return bestCampaign, true
	}
	sort.Sort(models.ByPrice(cs))
	if len(cs) > 0 {
		bestCampaign = cs[0]
	}
	return bestCampaign, false
}
