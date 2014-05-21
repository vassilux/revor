package mongo

import (
	"errors"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"revor/app/models"
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
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"call_daily":  1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
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
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"call_daily":  1,
		},
	}
	//
	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
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
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"call_daily":  1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
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
			"peer":        "$metadata.dst",
			"disposition": "$metadata.disposition",
			"call_daily":  1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
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

//month part
//Extract the numbers of calls by peer for given date
func GetMonthPeerInCalls(day string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlyanalytics_incomming")
	fmt.Printf("*********** GetMonthPeerInCalls enter\r\n")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
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
			"peer":         "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": 1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
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
	fmt.Printf("GetMonthPeerInCalls result %v for date %s\r\n", len(results), startDayDate)
	return results
}

//Extract the numbers of calls by peer for given date
func GetMonthPeerInCallsForUser(day string, user string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlyanalytics_incomming")
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
			"metadata.dst": user,
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":         "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": 1,
		},
	}
	//
	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$call_monthly"}},
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
	collection := mongoDb.C("monthlyanalytics_outgoing")
	fmt.Printf("*********** GetMonthPeerOutCalls enter\r\n")
	results := []bson.M{}
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), 1, 1, 0, 0, 0, time.UTC)
	//endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt": startDayDate, // bson.M{"$gte": startDayDate, "$lte": endDayDate},
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":         "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": 1,
		},
	}
	/*myGroup := bson.M{
		"$group": bson.M{
			"_id":        "$peer",
			"status":     bson.M{"$addToSet": "$disposition"},
			"callsCount": bson.M{"$addToSet": "$call_daily"},
		},
	}*/
	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
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
	collection := mongoDb.C("monthlyanalytics_outgoing")
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
			"metadata.dst": user,
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":         "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": 1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
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
	pipe := collection.Pipe(operations)
	err = pipe.All(&results)
	if err != nil {
		panic(err)
	}

	return results
}

//end of the month part
//Extract the numbers of calls by peer for given date
func GetYearPeerInCalls(year int, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlyanalytics_incomming")
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
			"peer":         "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": 1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
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

//Extract the numbers of calls by peer for given date
func GetYearPeerInCallsForUser(year int, user string, mongoDb *mgo.Database) []bson.M {
	incomming := mongoDb.C("monthlyanalytics_incomming")
	results := []bson.M{}
	startDayDate := time.Date(year, 1, 1, 1, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
			"metadata.dst": user,
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":         "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": 1,
		},
	}
	//
	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
			"dispositions": bson.M{"$addToSet": bson.M{"status": "$disposition", "callsCount": "$call_monthly"}},
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
	collection := mongoDb.C("monthlyanalytics_outgoing")
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
			"peer":         "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": 1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
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
	pipe := collection.Pipe(operations)
	err := pipe.All(&results)
	if err != nil {
		panic(err)
	}
	return results
}

//Extract the numbers of calls by peer for given date
func GetYearPeerOutCallsForUser(year int, user string, mongoDb *mgo.Database) []bson.M {
	collection := mongoDb.C("monthlyanalytics_outgoing")
	results := []bson.M{}

	startDayDate := time.Date(year, 1, 1, 1, 0, 0, 0, time.UTC)
	endDayDate := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	myMatch := bson.M{
		"$match": bson.M{
			"metadata.dt":  bson.M{"$gte": startDayDate, "$lte": endDayDate},
			"metadata.dst": user,
		},
	}
	//
	myProject := bson.M{
		"$project": bson.M{
			"peer":         "$metadata.dst",
			"disposition":  "$metadata.disposition",
			"call_monthly": 1,
		},
	}

	myGroup := bson.M{
		"$group": bson.M{
			"_id":          "$peer",
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

//CRUD part
func CreatePeer(id, value, comment string, mongoDb *mgo.Database) error {
	peer, err := GetPeer(id, mongoDb)
	if peer != nil {
		return err
	}
	var collection = mongoDb.C("peers")
	//
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
