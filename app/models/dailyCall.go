package models

import ()

type DailyCall struct {
	Id              string       `json:"id"					bson:"_id"`
	Meta            MetaData     `json:"metadata"			bson:"metadata"`
	AnswereWaitTime int          `json:"answerWaitTime"		bson:"answer_wait_time"`
	CallDaily       int          `json:"callDaily"			bson:"call_daily"`
	DurationDaily   int          `json:"durationDaily"		bson:"duration_daily"`
	CallsHourly     CallsByHours `json:"callHourly"		    bson:"call_hourly"`
}

type UserCalls struct {
	Id         string `json:"id"					bson:"_id"`
	CallsCount string `json:"callsCount"			bson:"calls_count"`
}
