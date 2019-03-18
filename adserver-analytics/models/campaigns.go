package models

import (
	"os"
)

// CampaignAnalytic struct
type CampaignAnalytic struct {
	Timestamp   string  `json:"timestamp,omitempty"`
	PlacementID string  `json:"placement"`
	CampaignID  string  `json:"campaign"`
	Price       float64 `json:"price"`
}

type CampaignManager struct {
	StoredData map[string]map[string]float64
	File       *os.File
	SumTotal   float64
}
