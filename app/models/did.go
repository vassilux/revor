package models

import (
	"fmt"
)

type Did struct {
	Did     string `bson:"did"			    json:"did"`
	Value   string `json:"value"			bson:"value"`
	Comment string `json:"comment"			bson:"comment"`
}

func (d *Did) String() string {
	return fmt.Sprintf("Did(%s)", d.Did)
}
