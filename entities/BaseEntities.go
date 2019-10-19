package entities

import (
	"time"
)

type TimelessModel struct {
	ID      string `gorm:"type:UUID;unique_index"`
}

type Model struct {
	ID      string `gorm:"type:UUID;unique_index"`
	Created time.Time
	Updated time.Time
	Deleted *time.Time
}

type ExpirableModel struct {
	ID        string `gorm:"type:UUID;unique_index"`
	Created   time.Time
	Deleted   *time.Time
	ValidFrom time.Time
	ValidTo   time.Time
}

type Nameable struct {
	Name string
}

type Aliasable struct {
	Alias string
}

type Enumerable struct {
	Number int64 `gorm:"AUTO_INCREMENT"`
}

type Authable struct {
	Password string `gorm:"type:CHAR(60)" json:"-"`
}

type Location struct {
	Latitude  float64
	Longitude float64
}

type Traceable struct {
	Country        string
	Ip             string
	Language       string
	IsBot          bool
	BrowserName    string
	BrowserVersion int64
	Platform       string
	OsName         string
	OsVersion      int64
	DeviceType     string
	Region         string
	City           string
	ZipCode        string
	Latitude       float64
	Longitude      float64
	Asn            string
	Referrer       string
	ReferrerHost   string
	Visits         int64
	IsConversion   bool
	IsNotified     bool
	Timestamp      time.Time `json:"timestamp,omitempty"`
}

type Wholesaleable struct {
	IsWholesale bool
}

type Priceable struct {
	Price       int64
	Tax         int64
	Points      int64
}

