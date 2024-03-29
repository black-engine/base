package entities

import "time"

type Request struct {
	ID string `gorm:"primary_key;type:uuid"`
	Created time.Time
	Method string
	Status int
	Path string
	Latency int64
}