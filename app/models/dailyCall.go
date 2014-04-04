package models

import ()

type DailyCall struct {
	Id              string
	MetaData        MetaData
	AnswereWaitTime int64
	CallDaily       int
	DurationDaily   int
}
