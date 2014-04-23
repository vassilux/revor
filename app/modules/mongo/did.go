package mongo

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

//Extract the numbers of calls by did for given date
func GetDidCalls(day string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailydid_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"did":        "$metadata.user",
			"call_daily": 1,
		},
	}
	myGroup := bson.M{
		"$group": bson.M{
			"_id":        "$did",
			"callsCount": bson.M{"$sum": "$call_daily"},
		},
	}
	mySort := bson.M{
		"$sort": bson.M{
			"callsCount": -1,
		},
	}
	//
	operations := []bson.M{myMatch, myProject, myGroup, mySort}
	pipe := incomming.Pipe(operations)
	err = pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}
