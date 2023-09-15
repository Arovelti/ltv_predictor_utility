package models

type UserData struct {
	DataType string
	Data     []LoadUserData
}

type LoadUserData struct {
	UserID     string `json:"UserId"`
	CampaignID string `json:"CampaignId"`
	Country    string `json:"Country"`
	Users      int    `json:"Users"`

	LTV1 float64 `json:"Ltv1"`
	LTV2 float64 `json:"Ltv2"`
	LTV3 float64 `json:"Ltv3"`
	LTV4 float64 `json:"Ltv4"`
	LTV5 float64 `json:"Ltv5"`
	LTV6 float64 `json:"Ltv6"`
	LTV7 float64 `json:"Ltv7"`
}

type AggregatedData struct {
	Key        string
	TotalLTV   float64
	UserCount  int
	AverageLtv float64
	MaxLtv     float64
	LTV60thDay float64
}
