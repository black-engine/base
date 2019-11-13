package entities

import "time"

type Campaign struct {
	Model
	Nameable
	DestinationUrl string
	FallbackUrl    string
}

type Prospect struct {
	Model
	Traceable
	Domain     string
	Campaign   *Campaign
	CampaignID *string `gorm:"type:UUID"`
}

type FlowEvent struct {
	ID      string `gorm:"type:UUID;unique_index"`
	Created time.Time
	Label   string

	Prospect   Prospect
	ProspectID string `gorm:"type:UUID"`
}
