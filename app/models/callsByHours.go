package models

import ()

type CallsByHours struct {
	H0 int `json:"00"		bson:"0"`
	H1 int `json:"01"		bson:"1"`
	H2 int `json:"02"		bson:"2"`
	H3 int `json:"03"		bson:"3"`
	H4 int `json:"04"		bson:"5"`
	H5 int `json:"05"		bson:"5"`
	H6 int `json:"06"		bson:"6"`
	H7 int `json:"07"		bson:"7"`
}
