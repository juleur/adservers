package utils

import "github.com/juleur/mini-adserver/adserver-analytics/models"

// NewCampaignManager s
func NewCampaignManager() *models.CampaignManager {
	cm := &models.CampaignManager{
		File:       TextOpener(),
		StoredData: make(map[string]map[string]float64),
	}
	return cm
}
