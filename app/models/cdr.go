package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Cdr struct {
	Id             bson.ObjectId `bson:"_id"`
	Calldate       time.Time     `bson:"call_date"`
	MetadataDt     time.Time     `bson:"metadata_date"`
	ClidName       string        `bson:"clid_name"`
	ClidNumber     string        `bson:"clid_number"`
	Src            string        `bson:"src"`
	Channel        string        `bson:"channel"`
	Dcontext       string        `bson:"dcontext"`
	DispositionStr string        `bson:"disposition_str"`
	Disposition    int           `bson:"disposition"`
	AnswerWaitTime int           `bson:"answer_wait_time"`
	Billsec        int           `bson:"billsec"`
	Duration       int           `bson:"duration"`
	Uniqueid       string        `bson:"uniqueid"`
	InoutStatus    int           `bson:"inout_status"`
	RecordFile     string        `bson:"record_file"`
	Dst            string        `bson:"dst"`
	Dnid           string        `bson:"dnid"`
	Dstchannel     string        `bson:"dst_channel"`
	CallDetails    []CallDetail  `bson:"call_details"`
}

type CallDetail struct {
	EventType string    `bson:"event_type"`
	EventTime time.Time `bson:"event_time"`
	CidNum    string    `bson:"cid_num"`
	CidDnid   string    `bson:"cid_dnid"`
	Exten     string    `bson:"exten"`
	UniqueId  string    `bson:"uniqueId"`
	LinkedId  string    `bson:"linkedId"`
	Peer      string    `bson:"peer"`
}
