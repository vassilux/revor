package mongo

import (
	"errors"
	"fmt"
	"github.com/jinzhu/now"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"revor/app/models"
	"time"
)

//Extract the numbers of calls by peer for given date
func GetPeerInCalls(day string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailypeer_incomming")
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
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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
func GetPeerInCallsByHours(day string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailypeer_incomming")
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
			"peer":            "$metadata.dst",
			"disposition":     "$metadata.disposition",
			"calls_per_hours": "$calls_per_hours",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callHourly": "$calls_per_hours"}},
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
func GetPeerOutCallsByHours(day string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailypeer_outgoing")
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
			"peer":            "$metadata.dst",
			"disposition":     "$metadata.disposition",
			"calls_per_hours": "$calls_per_hours",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callHourly": "$calls_per_hours"}},
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
func GetPeerInCallsByHoursAndPeer(day string, user string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailypeer_incomming")
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
			"metadata.dst": bson.RegEx{user, "i"},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":            "$metadata.dst",
			"disposition":     "$metadata.disposition",
			"calls_per_hours": "$calls_per_hours",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callHourly": "$calls_per_hours"}},
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
func GetPeerOutCallsByHoursAndPeer(day string, user string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailypeer_outgoing")
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
			"metadata.dst": bson.RegEx{user, "i"},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":            "$metadata.dst",
			"disposition":     "$metadata.disposition",
			"calls_per_hours": "$calls_per_hours",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callHourly": "$calls_per_hours"}},
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
	incomming := mongoDb.C("dailypeer_incomming")
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
			"metadata.dst": bson.RegEx{user, "i"},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}
	//
	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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
	collection := mongoDb.C("dailypeer_outgoing")
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
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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
	collection := mongoDb.C("dailypeer_outgoing")
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
			"metadata.dst": bson.RegEx{user, "i"},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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

//gen stats part

func GetPeerOutCallsAndDurationForDay(day string, peer string, mongoDb *mgo.Database) bson.M {
	collection := mongoDb.C("dailypeer_outgoing")
	results := bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst":         bson.RegEx{peer, "i"},
				"metadata.disposition": 16,
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.disposition": 16,
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"_id":      0,
			"calls":    "$calls",
			"duration": "$duration",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":                 0,
			"outCallsNumber":      bson.M{"$sum": "$calls"},
			"outCallsDuration":    bson.M{"$sum": "$duration"},
			"outCallsAvgDuration": bson.M{"$avg": "$duration"},
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup}
	pipe := collection.Pipe(operations)
	err = pipe.One(&results)
	if err != nil {
		results["outCallsNumber"] = 0
		results["outCallsDuration"] = 0
		results["outCallsAvgDuration"] = 0
		return results
	}

	return results
}

func GetPeerTotalInCallsForDay(day string, peer string, mongoDb *mgo.Database) bson.M {
	collection := mongoDb.C("dailypeer_incomming")
	results := bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"_id":   0,
			"calls": "$calls",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":     0,
			"inCalls": bson.M{"$sum": "$calls"},
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup}
	pipe := collection.Pipe(operations)
	err = pipe.One(&results)
	if err != nil {
		results["inCalls"] = 0
		return results
	}
	return results
}

func GetPeerInCallsAndDurationForDay(day string, peer string, mongoDb *mgo.Database) bson.M {
	collection := mongoDb.C("dailypeer_incomming")
	results := bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst":         bson.RegEx{peer, "i"},
				"metadata.disposition": 16,
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.disposition": 16,
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"_id":              0,
			"calls":            "$calls",
			"duration":         "$duration",
			"answer_wait_time": "$answer_wait_time",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":                      0,
			"inCallsAnswered":          bson.M{"$sum": "$calls"},
			"inCallsDuration":          bson.M{"$sum": "$duration"},
			"inCallsAvgDuration":       bson.M{"$avg": "$duration"},
			"inCallsAvgWaitAnswerTime": bson.M{"$avg": "$answer_wait_time"},
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup}
	pipe := collection.Pipe(operations)
	err = pipe.One(&results)
	if err != nil {
		results["inCallsAnswered"] = 0
		results["inCallsDuration"] = 0
		results["inCallsAvgDuration"] = 0
		results["inCallsAvgWaitAnswerTime"] = 0
		return results
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

//month part
//Extract the numbers of calls by peer for given date
func GetMonthPeerInCalls(day string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlypeer_incomming")

	results := []bson.M{}

	startDate, err := time.Parse(time.RFC3339, day)

	if err != nil {
		panic(err)
	}

	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)

	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt": startDayDate,
		},
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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
	fmt.Printf("GetMonthPeerInCalls result %v for date %s\r\n", len(results), startDayDate)
	return results
}

//Extract the numbers of calls by peer for given date
func GetMonthPeerInCallsForUser(day string, user string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlypeer_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)
	//endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt":  startDayDate, //bson.M{"$gte": startDayDate, "$lte": endDayDate},
			"metadata.dst": bson.RegEx{user, "i"},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}
	//
	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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
	fmt.Printf("GetMonthPeerInCallsForUser result %v \r\n", len(results))
	return results
}

//Extract the numbers of calls by peer for given date
func GetMonthPeerOutCalls(day string, mongoDb *mgo.Database) []bson.M {
	collection := mongoDb.C("monthlypeer_outgoing")

	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)
	//endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt": startDayDate,
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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
	fmt.Printf("GetMonthPeerOutCalls result %v for date %s\r\n", len(results), startDayDate)
	return results
}

//Extract the numbers of calls by peer for given date
func GetMonthPeerOutCallsForUser(day string, user string, mongoDb *mgo.Database) []bson.M {
	collection := mongoDb.C("monthlypeer_outgoing")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)
	//endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt":  startDayDate, //  bson.M{"$gte": startDayDate, "$lte": endDayDate},
			"metadata.dst": bson.RegEx{user, "i"},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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

//month part genstats
func GetMonthPeerOutCallsAndDurationForDay(day string, peer string, mongoDb *mgo.Database) bson.M {
	collection := mongoDb.C("monthlypeer_outgoing")
	results := bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          startDayDate,
				"metadata.dst":         bson.RegEx{peer, "i"},
				"metadata.disposition": 16,
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          startDayDate,
				"metadata.disposition": 16,
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"_id":      0,
			"calls":    "$calls",
			"duration": "$duration",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":                 0,
			"outCallsNumber":      bson.M{"$sum": "$calls"},
			"outCallsDuration":    bson.M{"$sum": "$duration"},
			"outCallsAvgDuration": bson.M{"$avg": "$duration"},
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup}
	pipe := collection.Pipe(operations)
	err = pipe.One(&results)
	if err != nil {
		results["outCallsNumber"] = 0
		results["outCallsDuration"] = 0
		results["outCallsAvgDuration"] = 0
		return results
	}

	return results
}

func GetMonthPeerTotalInCallsForDay(day string, peer string, mongoDb *mgo.Database) bson.M {
	collection := mongoDb.C("monthlypeer_incomming")
	results := bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}

	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  startDayDate,
				"metadata.dst": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": startDayDate,
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"_id":   0,
			"calls": "$calls",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":     0,
			"inCalls": bson.M{"$sum": "$calls"},
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup}
	pipe := collection.Pipe(operations)
	err = pipe.One(&results)
	if err != nil {
		results["inCalls"] = 0
		return results
	}
	return results
}

func GetMonthPeerInCallsAndDurationForDay(day string, peer string, mongoDb *mgo.Database) bson.M {
	collection := mongoDb.C("monthlypeer_incomming")
	results := bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}

	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          startDayDate,
				"metadata.dst":         bson.RegEx{peer, "i"},
				"metadata.disposition": 16,
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          startDayDate,
				"metadata.disposition": 16,
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"_id":              0,
			"calls":            "$calls",
			"duration":         "$duration",
			"answer_wait_time": "$answer_wait_time",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":                      0,
			"inCallsAnswered":          bson.M{"$sum": "$calls"},
			"inCallsDuration":          bson.M{"$sum": "$duration"},
			"inCallsAvgDuration":       bson.M{"$avg": "$duration"},
			"inCallsAvgWaitAnswerTime": bson.M{"$avg": "$answer_wait_time"},
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup}
	pipe := collection.Pipe(operations)
	err = pipe.One(&results)
	if err != nil {
		results["inCallsAnswered"] = 0
		results["inCallsDuration"] = 0
		results["inCallsAvgDuration"] = 0
		results["inCallsAvgWaitAnswerTime"] = 0
		return results
	}
	return results
}

//end of the month part
//Extract the numbers of calls by peer for given date
//year part
func GetYearPeerOutCallsAndDurationForYear(year int, peer string, mongoDb *mgo.Database) bson.M {
	collection := mongoDb.C("monthlypeer_outgoing")
	results := bson.M{}

	startDayDate := time.Date(year, 1, 1, 1, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst":         bson.RegEx{peer, "i"},
				"metadata.disposition": 16,
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.disposition": 16,
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"_id":      0,
			"calls":    "$calls",
			"duration": "$duration",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":                 0,
			"outCallsNumber":      bson.M{"$sum": "$calls"},
			"outCallsDuration":    bson.M{"$sum": "$duration"},
			"outCallsAvgDuration": bson.M{"$avg": "$duration"},
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup}
	pipe := collection.Pipe(operations)
	err := pipe.One(&results)
	if err != nil {
		results["outCallsNumber"] = 0
		results["outCallsDuration"] = 0
		results["outCallsAvgDuration"] = 0
		return results
	}

	return results
}

func GetYearPeerTotalInCallsForYear(year int, peer string, mongoDb *mgo.Database) bson.M {
	collection := mongoDb.C("monthlypeer_incomming")
	results := bson.M{}

	startDayDate := time.Date(year, 1, 1, 1, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"_id":   0,
			"calls": "$calls",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":     0,
			"inCalls": bson.M{"$sum": "$calls"},
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup}
	pipe := collection.Pipe(operations)
	err := pipe.One(&results)
	if err != nil {
		results["inCalls"] = 0
		return results
	}
	return results
}

func GetYearPeerInCallsAndDurationForYear(year int, peer string, mongoDb *mgo.Database) bson.M {
	collection := mongoDb.C("monthlypeer_incomming")

	results := bson.M{}

	startDayDate := time.Date(year, 1, 1, 1, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst":         bson.RegEx{peer, "i"},
				"metadata.disposition": 16,
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.disposition": 16,
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"_id":              0,
			"calls":            "$calls",
			"duration":         "$duration",
			"answer_wait_time": "$answer_wait_time",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":                      0,
			"inCallsAnswered":          bson.M{"$sum": "$calls"},
			"inCallsDuration":          bson.M{"$sum": "$duration"},
			"inCallsAvgDuration":       bson.M{"$avg": "$duration"},
			"inCallsAvgWaitAnswerTime": bson.M{"$avg": "$answer_wait_time"},
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup}
	pipe := collection.Pipe(operations)
	err := pipe.One(&results)
	if err != nil {
		results["inCallsAnswered"] = 0
		results["inCallsDuration"] = 0
		results["inCallsAvgDuration"] = 0
		results["inCallsAvgWaitAnswerTime"] = 0
		return results
	}
	return results
}

func GetYearPeerInCalls(year int, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlypeer_incomming")
	results := []bson.M{}

	startDayDate := time.Date(year, 1, 1, 1, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
		},
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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
	err := pipe.All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

//Extract the numbers of calls by peer for given date
func GetYearPeerInCallsForUser(year int, peer string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlypeer_incomming")
	results := []bson.M{}
	startDayDate := time.Date(year, 1, 1, 1, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst":         bson.RegEx{peer, "i"},
				"metadata.disposition": 16,
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.disposition": 16,
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}
	//
	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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
	err := pipe.All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

//Extract the numbers of calls by peer for given date
func GetYearPeerOutCalls(year int, mongoDb *mgo.Database) []bson.M {
	collection := mongoDb.C("monthlypeer_outgoing")
	results := []bson.M{}
	startDayDate := time.Date(year, 1, 1, 1, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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
	err := pipe.All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

//Extract the numbers of calls by peer for given date
func GetYearPeerOutCallsForUser(year int, user string, mongoDb *mgo.Database) []bson.M {
	collection := mongoDb.C("monthlypeer_outgoing")
	results := []bson.M{}

	startDayDate := time.Date(year, 1, 1, 1, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
			"metadata.dst": bson.RegEx{user, "i"},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$calls"}},
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
	err := pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

//
func GetPeer(id string, mongoDb *mgo.Database) (*models.Peer, error) {
	peer := models.Peer{}
	var collection = mongoDb.C("peers")
	var err = collection.Find(bson.M{"peer": id}).One(&peer)
	if err != nil {
		return nil, err
	}

	return &peer, nil
}

//To get peers datas calls/missing calls by month
//Incomming or outgoing datas depends of the parameter inout
func GetPeerYearInOutCallsByMonthAndPeer(year int, peer string, inout string, mongoDb *mgo.Database) []bson.M {
	var collectionName string
	results := []bson.M{}
	if inout == "in" {
		collectionName = "monthlypeerincomming_summary"
	} else if inout == "out" {
		collectionName = "monthlypeeroutgoing_summary"
	} else {
		return results
	}
	incomming := mongoDb.C(collectionName)

	startDayDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":   bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.peer": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
			},
		}
	}

	myProject := bson.M{
		"$project": bson.M{
			"_id":     bson.M{"$month": "$metadata.dt"},
			"calls":   "$calls",
			"missing": "$missing",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":     "$_id",
			"calls":   bson.M{"$sum": "$calls"},
			"missing": bson.M{"$sum": "$missing"},
		},
	}

	mySortByMonth := bson.M{
		"$sort": bson.M{
			"_id": 1,
		},
	}

	myProjectFinal := bson.M{
		"$project": bson.M{
			"_id":     0,
			"month":   "$_id",
			"calls":   1,
			"missing": 1,
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup, mySortByMonth, myProjectFinal}
	pipe := incomming.Pipe(operations)

	err := pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

//month part for peers
func doGetPeersMonthInDatasByDayAndPeer(day string, peer string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlypeer_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	//
	var myMatch bson.M
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)
	//check if the did is set and use it
	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  startDayDate,
				"metadata.dst": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": startDayDate,
			},
		}
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"disposition":        "$metadata.disposition",
			"duration_monthly":   "$duration_monthly",
			"calls_per_days":     "$calls_per_days",
			"durations_per_days": "$durations_daily",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":   "$disposition",
			"datas": bson.M{"$addToSet": bson.M{"callsDaily": "$calls_per_days", "durationsDaily": "$durations_per_days"}},
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

func GetPeersMonthInDatasByDay(day string, mongoDb *mgo.Database) []bson.M {
	return doGetPeersMonthInDatasByDayAndPeer(day, "", mongoDb)
}

func GetPeersMonthInDatasByDayAndPeer(day string, peer string, mongoDb *mgo.Database) []bson.M {
	return doGetPeersMonthInDatasByDayAndPeer(day, peer, mongoDb)
}

// outgoing peer datas
func doGetPeersMonthOutDatasByDayAndPeer(day string, peer string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlypeer_outgoing")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	//
	var myMatch bson.M
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)
	//check if the did is set and use it
	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  startDayDate,
				"metadata.dst": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": startDayDate,
			},
		}
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"disposition":        "$metadata.disposition",
			"duration_monthly":   "$duration_monthly",
			"calls_per_days":     "$calls_per_days",
			"durations_per_days": "$durations_per_days",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":   "$disposition",
			"datas": bson.M{"$addToSet": bson.M{"callsDaily": "$calls_per_days", "durationsDaily": "$durations_per_days"}},
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

func GetPeersMonthOutDatasByDay(day string, mongoDb *mgo.Database) []bson.M {
	return doGetPeersMonthOutDatasByDayAndPeer(day, "", mongoDb)
}

func GetPeersMonthOutDatasByDayAndPeer(day string, peer string, mongoDb *mgo.Database) []bson.M {
	return doGetPeersMonthOutDatasByDayAndPeer(day, peer, mongoDb)
}

func GetPeerInOutCallsDataByMonthForYearAndPeer(year int, peer string, inout string, mongoDb *mgo.Database) []bson.M {
	var collection string
	rawDatas := []bson.M{}
	if inout == "in" {
		collection = "monthlypeerincomming_summary"
	} else if inout == "out" {
		collection = "monthlypeeroutgoing_summary"
	} else {
		return rawDatas
	}
	targetCollection := mongoDb.C(collection)

	startDayDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":   bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.peer": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
			},
		}
	}

	myProject := bson.M{
		"$project": bson.M{
			"_id":      "$metadata.peer",
			"peer":     "$metadata.peer",
			"month":    bson.M{"$month": "$metadata.dt"},
			"calls":    "$calls",
			"missing":  "$missing",
			"duration": "$duration",
		},
	}

	mySortByMonth := bson.M{
		"$sort": bson.M{
			"peer":  1,
			"month": 1,
		},
	}

	percentProjectValue := bson.M{"$divide": []interface{}{"$missing", "$calls"}}
	percentProjectValuePercent := bson.M{"$multiply": []interface{}{"$percentRaw", 100}}
	myProjectPercent := bson.M{
		"$project": bson.M{
			"_id":        0,
			"peer":       1,
			"month":      1,
			"calls":      1,
			"missing":    1,
			"duration":   1,
			"percentRaw": percentProjectValue,
		},
	}

	myProjectPreFinal := bson.M{
		"$project": bson.M{
			"_id":        1,
			"peer":       1,
			"month":      1,
			"calls":      1,
			"missing":    1,
			"duration":   1,
			"percentRaw": 1,
			"percent":    percentProjectValuePercent,
		},
	}
	myProjectFinal := bson.M{
		"$project": bson.M{
			"_id":      0,
			"peer":     "$peer",
			"month":    "$month",
			"calls":    "$calls",
			"missing":  "$missing",
			"duration": "$duration",
			"percent":  "$percent",
		},
	}
	//
	operations := []bson.M{myMatch, myProject, mySortByMonth, myProjectPercent, myProjectPreFinal, myProjectFinal}
	pipe := targetCollection.Pipe(operations)

	err := pipe.All(&rawDatas)
	if err != nil {
		panic(err)
	}

	return rawDatas
}

func GetPeerInCallsDataByMonthForYearAndPeer(year int, peer string, mongoDb *mgo.Database) []bson.M {
	return GetPeerInOutCallsDataByMonthForYearAndPeer(year, peer, "in", mongoDb)
}

func GetPeerOutCallsDataByMonthForYearAndPeer(year int, peer string, mongoDb *mgo.Database) []bson.M {
	return GetPeerInOutCallsDataByMonthForYearAndPeer(year, peer, "out", mongoDb)
}

//To get general statistics for one day an if the peer number provide filter by the peer
func GetPeerGenStatsByDayAndPeer(day string, peer string, inout string, mongoDb *mgo.Database) []bson.M {
	var collectiionName string
	results := []bson.M{}
	if inout == "in" {
		collectiionName = "dailypeer_incomming"
	} else if inout == "out" {
		collectiionName = "dailypeer_outgoing"
	} else {
		return results
	}

	targetCollection := mongoDb.C(collectiionName)

	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":             "$metadata.dst",
			"disposition":      "$metadata.disposition",
			"calls":            "$calls",
			"duration":         "$duration",
			"answer_wait_time": "$answer_wait_time",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":              bson.M{"peer": "$peer", "disposition": "$disposition"},
			"calls":            bson.M{"$sum": "$calls"},
			"durations":        bson.M{"$sum": "$duration"},
			"answer_wait_time": bson.M{"$sum": "$answer_wait_time"},
		},
	}

	mySort := bson.M{
		"$sort": bson.M{
			"_id": 1,
		},
	}

	myProjectFinal := bson.M{
		"$project": bson.M{
			"_id":            0,
			"peer":           "$_id.peer",
			"disposition":    "$_id.disposition",
			"calls":          "$calls",
			"duration":       "$durations",
			"answerWaitTime": "$answer_wait_time",
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup, mySort, myProjectFinal}
	pipe := targetCollection.Pipe(operations)
	err = pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

//
func GetPeerWeekCallStatsByDayAndPeer(day string, peer string, inout string, mongoDb *mgo.Database) []models.HourByDaysRecord {
	var collectiionName string
	finalResults := make([]models.HourByDaysRecord, 24)
	if inout == "in" {
		collectiionName = "dailypeer_incomming"
	} else if inout == "out" {
		collectiionName = "dailypeer_outgoing"
	} else {
		return finalResults
	}

	incomming := mongoDb.C(collectiionName)
	results := []models.DayOfWeekCalls{}
	requestedDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	mondayDate := now.New(requestedDate).Monday()
	sundayDate := now.New(requestedDate).EndOfSunday()
	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": mondayDate, "$lte": sundayDate},
				"metadata.dst":         bson.RegEx{peer, "i"},
				"metadata.disposition": 16,
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":          bson.M{"$gte": mondayDate, "$lte": sundayDate},
				"metadata.disposition": 16,
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"day":              bson.M{"$dayOfWeek": "$metadata.dt"},
			"date":             "$metadata.dt",
			"disposition":      "$metadata.disposition",
			"calls":            "$calls",
			"duration":         "$duration",
			"answer_wait_time": "$answer_wait_time",
			"calls_per_hours":  "$calls_per_hours",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":              bson.M{"day": "$day", "disposition": "$disposition", "date": "$date"},
			"calls":            bson.M{"$sum": "$calls"},
			"durations":        bson.M{"$sum": "$duration"},
			"answer_wait_time": bson.M{"$sum": "$answer_wait_time"},
			"callsPerHours":    bson.M{"$addToSet": bson.M{"hourlyCalls": "$calls_per_hours"}},
		},
	}

	mySort := bson.M{
		"$sort": bson.M{
			"_id": 1,
		},
	}

	myProjectFinal := bson.M{
		"$project": bson.M{
			"_id":            0,
			"dayOfWeek":      "$_id.day",
			"disposition":    "$_id.disposition",
			"calls":          "$calls",
			"duration":       "$durations",
			"answerWaitTime": "$answer_wait_time",
			"date":           "$_id.date",
			"callsPerHours":  "$callsPerHours",
		},
	}
	//
	operations := []bson.M{myMatch, myProject, myGroup, mySort, myProjectFinal}
	pipe := incomming.Pipe(operations)
	err = pipe.All(&results)
	if err != nil {
		panic(err)
	}
	//

	for i := 0; i < 24; i++ {
		finalResults[i].Hour = fmt.Sprintf("%02d:00", i)
	}
	//
	for j := 0; j < len(results); j++ {
		//
		results[j].MakeSummaryCallsPerHours()
		results[j].PopulateHoursByDays(finalResults)
	}
	return finalResults
}

func GetPeerDispositionByDay(day string, inout string, peer string, mongoDb *mgo.Database) []bson.M {
	var collectiionName string
	results := []bson.M{}
	if inout == "in" {
		collectiionName = "dailypeer_incomming"
	} else if inout == "out" {
		collectiionName = "dailypeer_outgoing"
	} else {
		return results
	}

	targetCollection := mongoDb.C(collectiionName)

	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"disposition": "$metadata.disposition",
			"calls":       "$calls",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":   bson.M{"disposition": "$disposition"},
			"calls": bson.M{"$sum": "$calls"},
		},
	}

	mySort := bson.M{
		"$sort": bson.M{
			"_id": 1,
		},
	}

	myProjectFinal := bson.M{
		"$project": bson.M{
			"_id":         0,
			"disposition": "$_id.disposition",
			"calls":       "$calls",
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup, mySort, myProjectFinal}
	pipe := targetCollection.Pipe(operations)
	err = pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

func GetPeerDispositionByMonth(day string, inout string, peer string, mongoDb *mgo.Database) []bson.M {
	var collectiionName string
	results := []bson.M{}
	if inout == "in" {
		collectiionName = "monthlypeer_incomming"
	} else if inout == "out" {
		collectiionName = "monthlypeer_outgoing"
	} else {
		return results
	}

	targetCollection := mongoDb.C(collectiionName)

	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}

	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  startDayDate,
				"metadata.dst": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": startDayDate,
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"disposition": "$metadata.disposition",
			"calls":       "$calls",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":   bson.M{"disposition": "$disposition"},
			"calls": bson.M{"$sum": "$calls"},
		},
	}

	mySort := bson.M{
		"$sort": bson.M{
			"_id": 1,
		},
	}

	myProjectFinal := bson.M{
		"$project": bson.M{
			"_id":         0,
			"disposition": "$_id.disposition",
			"calls":       "$calls",
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup, mySort, myProjectFinal}
	pipe := targetCollection.Pipe(operations)
	err = pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

func GetPeerDispositionByYear(year int, inout string, peer string, mongoDb *mgo.Database) []bson.M {
	var collectiionName string
	results := []bson.M{}
	if inout == "in" {
		collectiionName = "monthlypeer_incomming"
	} else if inout == "out" {
		collectiionName = "monthlypeer_outgoing"
	} else {
		return results
	}

	targetCollection := mongoDb.C(collectiionName)

	startDayDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(peer) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst": bson.RegEx{peer, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"disposition": "$metadata.disposition",
			"calls":       "$calls",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":   bson.M{"disposition": "$disposition"},
			"calls": bson.M{"$sum": "$calls"},
		},
	}

	mySort := bson.M{
		"$sort": bson.M{
			"_id": 1,
		},
	}

	myProjectFinal := bson.M{
		"$project": bson.M{
			"_id":         0,
			"disposition": "$_id.disposition",
			"calls":       "$calls",
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup, mySort, myProjectFinal}
	pipe := targetCollection.Pipe(operations)
	err := pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

//CRUD part
func CreatePeer(id, value, comment string, mongoDb *mgo.Database) error {
	peer, err := GetPeer(id, mongoDb)
	if peer != nil {
		return err
	}
	var collection = mongoDb.C("peers")
	//
	newPeer := models.Peer{}
	newPeer.Peer = id
	newPeer.Value = value
	newPeer.Comment = comment
	err = collection.Insert(&newPeer)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePeer(id, value, comment string, mongoDb *mgo.Database) error {
	peer := models.Peer{}
	var collection = mongoDb.C("peers")
	change := mgo.Change{
		Update:    bson.M{"peer": id, "value": value, "comment": comment},
		ReturnNew: true,
	}
	_, err := collection.Find(bson.M{"peer": id}).Apply(change, &peer)
	if err != nil {
		return err
	}
	if peer.Peer != id {
		msg := fmt.Sprint("Can't udpate a peer with the given id %s.", id)
		return errors.New(msg)
	}
	return nil
}

func DeletePeer(id string, mongoDb *mgo.Database) error {

	peer, err := GetPeer(id, mongoDb)
	if peer == nil {
		return err
	}
	var collection = mongoDb.C("peers")
	//
	var selector = bson.M{"peer": id}
	err = collection.Remove(selector)
	if err != nil {
		return err
	}

	return nil
}
