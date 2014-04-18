package models

import (
	"time"
)

type MetaData struct {
	User        string    `json:"user" 			bson:"user"`
	Dt          time.Time `json:"date"			bson:"dt"`
	Disposition int       `json:"disposition"	bson:"disposition"`
}
