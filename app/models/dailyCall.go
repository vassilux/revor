package models

import ()

type DailyCall struct {
	Id              string   `bson:"_id"`
	Meta            MetaData `bson:"metadata"`
	AnswereWaitTime int      `bson:"answer_wait_time"`
	CallDaily       int      `bson:"call_daily"`
	DurationDaily   int      `bson:"duration_daily"`
}
