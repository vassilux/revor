package models

import (
	"time"
)

type MetaData struct {
	User        string    `bson:"user"`
	Dt          time.Time `bson:"dt"`
	Disposition int       `bson:"disposition"`
}
