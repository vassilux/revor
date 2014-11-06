package mongo

import (
	"errors"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"revor/app/models"
	"time"
)

//helper for get the did calls for a day and the did if the provided
func doGetDidCallsByDayAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailydid_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(did) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst": bson.RegEx{did, "i"},
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
			"did":         "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       "$calls",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
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

//Extract the numbers of calls by did for given date
func GetDidCalls(day string, mongoDb *mgo.Database) []bson.M {
	return doGetDidCallsByDayAndDid(day, "", mongoDb)
}

//Extract the numbers of calls of the given did and the given given date
func GetDidCallsByDayAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
	return doGetDidCallsByDayAndDid(day, did, mongoDb)
}

func doGetDidCallsByHoursAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailydid_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	var myMatch bson.M

	if len(did) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst": bson.RegEx{did, "i"},
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
			"did":             "$metadata.dst",
			"disposition":     "$metadata.disposition",
			"calls_per_hours": "$calls_per_hours",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
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

//Extract the numbers of calls by did for given date
func GetDidCallsByHours(day string, mongoDb *mgo.Database) []bson.M {
	return doGetDidCallsByHoursAndDid(day, "", mongoDb)
}

//Extract the numbers of calls for did for given date
func GetDidCallsByHoursAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
	return doGetDidCallsByHoursAndDid(day, did, mongoDb)
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
func doGetDidCallsByMonthAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlydid_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)
	var myMatch bson.M

	if len(did) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  startDayDate, //bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst": bson.RegEx{did, "i"},
			},
		}
	} else {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt": startDayDate, //bson.M{"$gte": startDayDate, "$lte": endDayDate},
			},
		}
	}

	//
	myProject := bson.M{
		"$project": bson.M{
			"did":         "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       "$calls",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
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

//Extract the numbers of calls of the given did and the given given date
func GetDidCallsByMonthAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
	return doGetDidCallsByMonthAndDid(day, did, mongoDb)
}

//Extract the numbers of calls by did for given date
func GetDidMonthCalls(day string, mongoDb *mgo.Database) []bson.M {
	return doGetDidCallsByMonthAndDid(day, "", mongoDb)
}

//helper for did calls and duration for the given month by day of the month and filtred by the did if provided
func doGetMonthCallsByDayAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlydid_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	//
	var myMatch bson.M
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)
	//check if the did is set and use it
	if len(did) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  startDayDate,
				"metadata.dst": bson.RegEx{did, "i"},
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
			"disposition":      "$metadata.disposition",
			"duration_monthly": "$duration_monthly",
			"calls":            "$calls_per_days",
			"durations":        "$durations_per_days",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":   "$disposition",
			"datas": bson.M{"$addToSet": bson.M{"callsDaily": "$calls", "durationsDaily": "$durations"}},
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

//Extract the numbers of calls by did for the given date by days of the month
func GetDidsMonthCallsByDay(day string, mongoDb *mgo.Database) []bson.M {
	//call helper with the empty did
	return doGetMonthCallsByDayAndDid(day, "", mongoDb)
}

//Extract the numbers of calls by did for the given date by days of the month and filter by did
func GetDidsMonthCallsByDayAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
	//call helper
	return doGetMonthCallsByDayAndDid(day, did, mongoDb)
}

//yearly part
//helper
func doGetDidCallsByYearAndDid(year int, did string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlydid_incomming")
	results := []bson.M{}
	//
	startDayDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(did) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst": bson.RegEx{did, "i"},
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
			"did":         "$metadata.dst",
			"disposition": "$metadata.disposition",
			"calls":       "$calls",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$did",
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

//Extract the numbers of calls of the given did and the given given date
func GetDidCallsByYearAndDid(year int, did string, mongoDb *mgo.Database) []bson.M {
	return doGetDidCallsByYearAndDid(year, did, mongoDb)
}

//Extract the numbers of calls by did for given date
func GetDidYearCalls(year int, mongoDb *mgo.Database) []bson.M {
	return doGetDidCallsByYearAndDid(year, "", mongoDb)
}

//helper to get dids data for a given year by month of the year and filtred if the did provided
func doGetDidYearCallsByMonthAndDid(year int, did string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlydid_summary")
	rawDatas := []bson.M{}

	startDayDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(did) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.did": bson.RegEx{did, "i"},
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

	err := pipe.All(&rawDatas)
	if err != nil {
		panic(err)
	}

	return rawDatas

}

//To get dids all calls, missing calls and duration
//Also provide the percent of missing calls
func GetDidCallsDataByMonthForYearAndDid(year int, did string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlydid_summary")
	rawDatas := []bson.M{}

	startDayDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)

	var myMatch bson.M

	if len(did) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.did": bson.RegEx{did, "i"},
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
			"_id":      "$metadata.did",
			"did":      "$metadata.did",
			"month":    bson.M{"$month": "$metadata.dt"},
			"calls":    "$calls",
			"missing":  "$missing",
			"duration": "$duration",
		},
	}

	mySortByMonth := bson.M{
		"$sort": bson.M{
			"did":   1,
			"month": 1,
		},
	}

	percentProjectValue := bson.M{"$divide": []interface{}{"$missing", "$calls"}}
	percentProjectValuePercent := bson.M{"$multiply": []interface{}{"$percentRaw", 100}}
	myProjectPercent := bson.M{
		"$project": bson.M{
			"_id":        0,
			"did":        1,
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
			"did":        1,
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
			"did":      "$did",
			"month":    "$month",
			"calls":    "$calls",
			"missing":  "$missing",
			"duration": "$duration",
			"percent":  "$percent",
		},
	}
	//
	operations := []bson.M{myMatch, myProject, mySortByMonth, myProjectPercent, myProjectPreFinal, myProjectFinal}
	pipe := incomming.Pipe(operations)

	err := pipe.All(&rawDatas)
	if err != nil {
		panic(err)
	}

	return rawDatas
}

//To get general statistics for one day an if the did number provide filter by the did
func GetDidGenStatsByDayAndDid(day string, did string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("dailydid_incomming")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	var myMatch bson.M

	if len(did) > 0 {
		myMatch = bson.M{
			"$match": bson.M{
				"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
				"metadata.dst": bson.RegEx{did, "i"},
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
			"did":              "$metadata.dst",
			"disposition":      "$metadata.disposition",
			"calls":            "$calls",
			"duration":         "$duration",
			"answer_wait_time": "$answer_wait_time",
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":              bson.M{"did": "$did", "disposition": "$disposition"},
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
			"did":            "$_id.did",
			"disposition":    "$_id.disposition",
			"calls":          "$calls",
			"duration":       "$durations",
			"answerWaitTime": "$answer_wait_time",
		},
	}

	//
	operations := []bson.M{myMatch, myProject, myGroup, mySort, myProjectFinal}
	pipe := incomming.Pipe(operations)
	err = pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

//Extract the numbers of calls by did for given date
func GetDidYearCallsByMonth(year int, mongoDb *mgo.Database) []bson.M {
	return doGetDidYearCallsByMonthAndDid(year, "", mongoDb)
}

//Extract the numbers of calls by did for given date
func GetDidYearCallsByMonthAndDid(year int, did string, mongoDb *mgo.Database) []bson.M {
	return doGetDidYearCallsByMonthAndDid(year, did, mongoDb)
}

//
func GetDid(id string, mongoDb *mgo.Database) (*models.Did, error) {
	did := models.Did{}
	var collection = mongoDb.C("dids")
	var err = collection.Find(bson.M{"did": id}).One(&did)
	if err != nil {
		return nil, err
	}

	return &did, nil
}

//CRUD part
func CreateDid(id, value, comment string, mongoDb *mgo.Database) error {
	did, err := GetDid(id, mongoDb)
	if did != nil {
		return err
	}
	var collection = mongoDb.C("dids")
	//
	//
	newDid := models.Did{}
	newDid.Did = id
	newDid.Value = value
	newDid.Comment = comment
	err = collection.Insert(&newDid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDid(id, value, comment string, mongoDb *mgo.Database) error {
	did := models.Did{}
	var collection = mongoDb.C("dids")
	change := mgo.Change{
		Update:    bson.M{"did": id, "value": value, "comment": comment},
		ReturnNew: true,
	}
	_, err := collection.Find(bson.M{"did": id}).Apply(change, &did)
	if err != nil {
		return err
	}
	if did.Did != id {
		msg := fmt.Sprint("Can't udpate a did with the given id %s.", id)
		return errors.New(msg)
	}
	return nil
}

func DeleteDid(id string, mongoDb *mgo.Database) error {

	did, err := GetDid(id, mongoDb)
	if did == nil {
		return err
	}
	var collection = mongoDb.C("dids")
	//
	var selector = bson.M{"did": id}
	err = collection.Remove(selector)
	if err != nil {
		return err
	}

	return nil
}
