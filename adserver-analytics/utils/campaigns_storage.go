package utils

import (
	"bufio"
	"encoding/json"

	"github.com/juleur/mini-adserver/adserver-analytics/models"
)

// StoreCampaigns s
func StoreCampaigns(campaign models.CampaignAnalytic, storedCampaigns map[string]map[string]float64) {
	_, ok := storedCampaigns[campaign.PlacementID]
	if ok {
		storedCampaigns[campaign.PlacementID][campaign.CampaignID] += campaign.Price
	} else {
		storedCampaigns[campaign.PlacementID] = make(map[string]float64)
		storedCampaigns[campaign.PlacementID][campaign.CampaignID] = campaign.Price
	}
}

// StoreDataAtBoot s
func StoreDataAtBoot(cm *models.CampaignManager) {
	scanner := bufio.NewScanner(cm.File)
	for scanner.Scan() {
		c := models.CampaignAnalytic{}
		err := json.Unmarshal([]byte(scanner.Text()), &c)
		CheckError(err)
		StoreCampaigns(c, cm.StoredData)
		cm.SumTotal += c.Price
	}
}
