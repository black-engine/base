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
	CountryId string `gorm:"type:UUID"`
}

type City struct {
	TimelessModel
	Nameable
	Location

	ShippingCost int32
	ShippingTax  int32

	State   State
	StateId string `gorm:"type:UUID"`
}

type Neighborhood struct {
	TimelessModel
	Nameable
	Location

	Zipcode string
	ShippingCost int32
	ShippingTax  int32

	City   City
	CityId string `gorm:"type:UUID"`
}

type Branch struct {
	Model
	Aliasable
	Nameable
	Location
}

type BelongsToBranch struct {
	Branch   *Branch
	BranchId *string `gorm:"type:UUID"`
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
