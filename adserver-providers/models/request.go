package models

// BodyRequest struct
type BodyRequest struct {
	Country     string `json:"country"`
	Device      string `json:"device"`
	PlacementID string `json:"placement"`
}
