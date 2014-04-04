package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Cdr struct {
	Id             bson.ObjectId `bson:"_id"`
	Calldate       time.Time     `bson:"calldate"`
	MetadataDt     time.Time     `bson:"metadataDt"`
	ClidName       string        `bson:"clidName"`
	ClidNumber     string        `bson:"clidNumber"`
	Src            string        `bson:"src"`
	Channel        string        `bson:"channel"`
	Dcontext       string        `bson:"dcontext"`
	Disposition    int           `bson:"disposition"`
	Answerwaittime int64         `bson:"answerwaittime"`
	Billsec        int           `bson:"billsec"`
	Duration       int           `bson:"duration"`
	Uniqueid       string        `bson:"uniqueid"`
	Inoutstatus    int           `bson:"inoutstatus"`
	Recordfile     string        `bson:"recordfile"`
	Dst            string        `bson:"dst"`
}
