package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Cdr struct {
	Id             bson.ObjectId `json:"id"                 bson:"_id"`
	Calldate       time.Time     `json:"callDate"			bson:"call_date"`
	MetadataDt     time.Time     `json:"metadataDate"		bson:"metadata_date"`
	ClidName       string        `json:"clidName"           bson:"clid_name"`
	ClidNumber     string        `json:"clidNumber"		    bson:"clid_number"`
	Src            string        `json:"src"				bson:"src"`
	Channel        string        `json:"channel" 			bson:"channel"`
	Dcontext       string        `json:"dcontext"			bson:"dcontext"`
	DispositionStr string        `json:"dispositionStr"		bson:"disposition_str"`
	Disposition    int           `json:"disposition"		bson:"disposition"`
	AnswerWaitTime int           `json:"answerWaitTime"		bson:"answer_wait_time"`
	Billsec        int           `json:"billSec"			bson:"billsec"`
	Duration       int           `json:"duration"			bson:"duration"`
	UniqueId       string        `json:"uniqueId"			bson:"uniqueid"`
	InoutStatus    int           `json:"inoutStatus"		bson:"inout_status"`
	RecordFile     string        `json:"recordFile"			bson:"record_file"`
	Dst            string        `json:"dst"				bson:"dst"`
	Dnid           string        `json:"dnid"				bson:"dnid"`
	Dstchannel     string        `json:"dstChannel"			bson:"dst_channel"`
	CallDetails    []CallDetail  `json:"callDetails"		bson:"call_details"`
}

type CallDetail struct {
	EventType string    `json:"eventType"		bson:"event_type"`
	EventTime time.Time `json:"eventTime"		bson:"event_time"`
	CidNum    string    `json:"cidNum"			bson:"cid_num"`
	CidDnid   string    `json:"cidDnid"			bson:"cid_dnid"`
	Exten     string    `json:"exten"			bson:"exten"`
	UniqueId  string    `json:"uniqueId"		bson:"uniqueId"`
	LinkedId  string    `json:"linkedId"		bson:"linkedId"`
	Peer      string    `json:"peer"			bson:"peer"`
}
