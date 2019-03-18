package models

// ByPrice sort.interface for []Campaign based on Price field
type ByPrice []Campaign

func (p ByPrice) Len() int           { return len(p) }
func (p ByPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByPrice) Less(i, j int) bool { return p[i].Price > p[j].Price }

// Campaigns map of Campaign struct
type Campaigns struct {
	Campaigns map[string]Campaign
}

// Campaign struct
type Campaign struct {
	ID         string   `json:"campaign,omitempty"`
	Price      float64  `json:"price,omitempty"`
	Content    Content  `json:"content,omitempty"`
	Countries  []string `json:"countries,omitempty"`
	Devices    []string `json:"devices,omitempty"`
	Placements []string `json:"placements,omitempty"`
}

// Content struct
type Content struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Landing     string `json:"landing,omitempty"`
}

// CampaignAnalytic struct
type CampaignAnalytic struct {
	Timestamp   string  `json:"timestamp,omitempty"`
	PlacementID string  `json:"placement"`
	CampaignID  string  `json:"campaign"`
	Price       float64 `json:"price"`
}
