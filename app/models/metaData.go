package models

import (
	"time"
)

type MetaData struct {
	User        string
	Dt          time.Time
	Disposition int
}
