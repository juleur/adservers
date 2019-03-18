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
func GetCampaign(bodyReq models.BodyRequest, jsonData []models.Campaign) models.Campaign {
	cs := []models.Campaign{}
	bestCampaign := models.Campaign{}
	for _, campaignData := range jsonData {
		keep := isPlacementOk(bodyReq.PlacementID, campaignData.Placements)
		keep = isCountryOk(bodyReq.Country, campaignData.Countries)
		keep = isDeviceOk(bodyReq.Device, campaignData.Devices)
		if keep {
			campaignData.Placements[0] = bodyReq.PlacementID
			cs = append(cs, campaignData)
		}
	}
	sort.Sort(models.ByPrice(cs))
	if len(cs) > 0 {
		bestCampaign = cs[0]
	}
	return bestCampaign
}
