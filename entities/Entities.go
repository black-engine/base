package entities

import (
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Country struct {
	TimelessModel
	Nameable
	Aliasable
	Location
}

type State struct {
	TimelessModel
	Nameable
	Aliasable
	Location

	Country   Country
	CountryID string `gorm:"type:UUID"`
}

type City struct {
	TimelessModel
	Nameable
	Location

	ShippingCost int32
	ShippingTax  int32

	State   State
	StateID string `gorm:"type:UUID"`
}

type Neighborhood struct {
	TimelessModel
	Nameable
	Location

	Zipcode string
	ShippingCost int32
	ShippingTax  int32

	City   City
	CityID string `gorm:"type:UUID"`
}

type Branch struct {
	Model
	Aliasable
	Nameable
	Location
}

type BelongsToBranch struct {
	Branch   *Branch
	BranchID *string `gorm:"type:UUID"`
}

type Facet struct {
	TimelessModel
	Key string

	IsContinuous bool
	LowerBound   *float64
	UpperBound   *float64

	IsDiscrete  bool
	IsNumerical bool
	Values      *postgres.Jsonb
	Unit        *string
}

type BaseImage struct {
	ExpirableModel

	Position    int8
	Url         string
	IsThumbnail bool
	IsDefault   bool
}
