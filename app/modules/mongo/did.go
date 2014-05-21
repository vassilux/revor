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
			"did":         "$metadata.dst",
			"disposition": "$metadata.disposition",
			"call_daily":  "$call_daily",
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
	operations := []bson.M{myMatch, myProject, myGroup, mySort}
	pipe := incomming.Pipe(operations)
	err = pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

//Extract the numbers of calls of the given did and the given given date
func GetDidCallsByDayAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
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
			"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
			"metadata.dst": did,
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"did":         "$metadata.dst",
			"disposition": "$metadata.disposition",
			"call_daily":  "$call_daily",
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
	operations := []bson.M{myMatch, myProject, myGroup, mySort}
	pipe := incomming.Pipe(operations)
	err = pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

func GetDids(mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dids")
	results := []bson.M{}

	//
	myProject := bson.M{
		"$project": bson.M{
			"id":      "$did",
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
//Extract the numbers of calls of the given did and the given given date
func GetDidCallsByMonthAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlydid_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)
	//endDayDate := time.Date(startDate.Year(), startDate.Month(), 31, 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt":  startDayDate, //bson.M{"$gte": startDayDate, "$lte": endDayDate},
			"metadata.dst": did,
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"did":          "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": "$call_monthly",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$call_monthly"}},
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

//Extract the numbers of calls by did for given date
func GetDidMonthCalls(day string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlydid_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	//
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)
	//endDayDate := time.Date(startDate.Year(), startDate.Month(), 31, 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt": startDayDate, //bson.M{"$gte": startDayDate, "$lte": endDayDate},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"did":          "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": "$call_monthly",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$call_monthly"}},
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

//yearly part
//Extract the numbers of calls of the given did and the given given date
func GetDidCallsByYearAndDid(year int, did string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlydid_incomming")
	results := []bson.M{}
	//
	startDayDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
			"metadata.dst": did,
		},
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"did":          "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": "$call_monthly",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$call_monthly"}},
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

//Extract the numbers of calls by did for given date
func GetDidYearCalls(year int, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlydid_incomming")
	results := []bson.M{}
	startDayDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"did":          "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": "$call_monthly",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$call_monthly"}},
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
