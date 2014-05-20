package mongo

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

//Extract the numbers of calls by peer for given date
func GetPeerInCalls(day string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailyanalytics_incomming")
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
			"did":         "$metadata.dst",
			"disposition": "$metadata.disposition",
			"call_daily":  1,
		},
	}
	/*myGroup := bson.M{
		"$group": bson.M{
			"_id":        "$did",
			"status":     bson.M{"$addToSet": "$disposition"},
			"callsCount": bson.M{"$addToSet": "$call_daily"},
		},
	}*/
	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$call_daily"}},
		},
	}

	mySort := bson.M{
		"$sort": bson.M{
			"_id": 1,
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

//Extract the numbers of calls by peer for given date
func GetPeerInCallsForUser(day string, user string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailyanalytics_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
			"metadata.dst": user,
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"did":         "$metadata.dst",
			"disposition": "$metadata.disposition",
			"call_daily":  1,
		},
	}
	//
	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$call_daily"}},
		},
	}
	//
	mySort := bson.M{
		"$sort": bson.M{
			"_id": 1,
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

//Extract the numbers of calls by peer for given date
func GetPeerOutCalls(day string, mongoDb *mgo.Database) []bson.M {
	collection := mongoDb.C("dailyanalytics_outgoing")
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
			"did":         "$metadata.dst",
			"disposition": "$metadata.disposition",
			"call_daily":  1,
		},
	}
	/*myGroup := bson.M{
		"$group": bson.M{
			"_id":        "$did",
			"status":     bson.M{"$addToSet": "$disposition"},
			"callsCount": bson.M{"$addToSet": "$call_daily"},
		},
	}*/
	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$call_daily"}},
		},
	}

	mySort := bson.M{
		"$sort": bson.M{
			"_id": 1,
		},
	}
	//
	operations := []bson.M{myMatch, myProject, myGroup, mySort}
	pipe := collection.Pipe(operations)
	err = pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

//Extract the numbers of calls by peer for given date
func GetPeerOutCallsForUser(day string, user string, mongoDb *mgo.Database) []bson.M {
	collection := mongoDb.C("dailyanalytics_outgoing")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
			"metadata.dst": user,
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"did":         "$metadata.dst",
			"disposition": "$metadata.disposition",
			"call_daily":  1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$call_daily"}},
		},
	}

	mySort := bson.M{
		"$sort": bson.M{
			"_id": 1,
		},
	}
	//
	fmt.Printf("GetPeerOutCallsForUser myMatch %v \r\n", myMatch)
	//
	operations := []bson.M{myMatch, myProject, myGroup, mySort}
	pipe := collection.Pipe(operations)
	err = pipe.All(&results)
	fmt.Printf(" GetPeerOutCallsForUser %v \r\n", results)
	if err != nil {
		panic(err)
	}

	return results
}

func GetPeers(mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("peers")
	results := []bson.M{}

	//
	myProject := bson.M{
		"$project": bson.M{
			"id":      "$peer",
			"comment": "$comment",
			"value":   "$value",
		},
	}

	mySort := bson.M{
		"$sort": bson.M{
			"id": 1,
		},
	}
	//
	operations := []bson.M{myProject, mySort}
	pipe := incomming.Pipe(operations)
	err := pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}
