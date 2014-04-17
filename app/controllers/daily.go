package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"revor/app/models"
	"revor/app/modules/mongo"
	"time"
)

type Daily struct {
	App
}

//Fetch all incomming calls
//return a json stream with calls on success otherwise http status 500
func (c Daily) IncommingCalls() revel.Result {
	incomming := c.MongoDatabase.C("dailyanalytics_incomming")
	var startDate = time.Now()
	var searchResults []models.DailyCall
	var query = bson.M{"metadata.dt": startDate}
	var err = incomming.Find(query).Limit(mongo.DB_REQUEST_LIMITS).All(&searchResults)
	if err != nil {
		revel.ERROR.Printf("[Dayly] Got error from mongo : [%v].\r\n", err)
		panic(err)
	}
	return c.RenderJson(searchResults)

}

//Fetch all incomming calls for given day
//return a json stream with calls on success otherwise http status 500
func (c Daily) IncommingCallsByDay(day string) revel.Result {
	revel.TRACE.Printf("[Daily] Get incomming call by day for the date [%s].\r\n", day)
	incomming := c.MongoDatabase.C("dailyanalytics_incomming")
	var searchResults []models.DailyCall
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		revel.WARN.Printf("[Daily] Failed to parse the start date : [%v].\r\n", err)
		return c.RenderJson(searchResults)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	var query = bson.M{"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate}}
	err = incomming.Find(query).Limit(mongo.DB_REQUEST_LIMITS).All(&searchResults)
	if err != nil {
		revel.WARN.Printf("[Daily] Got error from mongo : [%v].\r\n", err)
		panic(err)
	}
	revel.TRACE.Printf("[Daily] send to the client response of %d records.\r\n", len(searchResults))
	return c.RenderJson(searchResults)
}

//Fetch all incomming calls for given day and user
//return a json stream with calls on success otherwise http status 500
func (c Daily) IncommingCallsByDayUser(day string, user string) revel.Result {
	revel.TRACE.Printf("[Daily] Get incomming call by day for the date [%s] and caller [%s].\r\n", day, user)
	incomming := c.MongoDatabase.C("dailyanalytics_incomming")
	var searchResults []models.DailyCall
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		revel.WARN.Printf("[Daily] Failed to parse the start date : [%v].\r\n", err)
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	var query = bson.M{"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate}, "metadata.user": user}
	err = incomming.Find(query).Limit(mongo.DB_REQUEST_LIMITS).All(&searchResults)
	if err != nil {
		revel.WARN.Printf("[Daily] Got error from mongo : [%v].\r\n", err)
		return c.RenderJson(searchResults)
	}
	revel.TRACE.Printf("[Daily] send to the client response of %d records.\r\n", len(searchResults))
	return c.RenderJson(searchResults)
}

//did parts
//Fetch all incomming calls for given day and caller
//return a json stream with calls on success otherwise http status 500
func (c Daily) IncommingDidCallsByDay(day string) revel.Result {
	revel.TRACE.Printf("[Daily DID] Get incomming call by day for the date [%s] and did [%s].\r\n", day)
	incomming := c.MongoDatabase.C("dailydid_incomming")
	var searchResults []models.DailyCall
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		revel.WARN.Printf("[Daily DID] Failed to parse the start date : [%v].\r\n", err)
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	revel.TRACE.Printf("[Daily DID] the start date [%s] and end date [%s].\r\n", startDayDate, endDayDate)
	//
	var query = bson.M{"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate}}

	err = incomming.Find(query).Limit(mongo.DB_REQUEST_LIMITS).All(&searchResults)
	if err != nil {
		revel.WARN.Printf("[Daily DID] Got error from mongo : [%v].\r\n", err)
		return c.RenderJson(searchResults)
	}
	revel.TRACE.Printf("[Daily DID] send to the client response of %d records.\r\n", len(searchResults))
	return c.RenderJson(searchResults)
}

//
func (c Daily) IncommingDidCallsByDayDid(day string, did string) revel.Result {
	revel.TRACE.Printf("[Daily DID] Get incomming call by day for the date [%s] and did [%s].\r\n", day, did)
	incomming := c.MongoDatabase.C("dailydid_incomming")
	var searchResults []models.DailyCall
	startDate, err := time.Parse(time.RFC3339, day)
	if err != nil {
		revel.WARN.Printf("[Daily DID] Failed to parse the start date : [%v].\r\n", err)
		panic(err)
	}
	startDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, time.UTC)
	endDayDate := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 23, 59, 59, 0, time.UTC)
	revel.TRACE.Printf("[Daily DID] the start date [%s] and end date [%s].\r\n", startDayDate, endDayDate)
	//
	var query = bson.M{"metadata.dt": bson.M{"$gte": startDayDate, "$lte": endDayDate}, "metadata.user": did}

	err = incomming.Find(query).Limit(mongo.DB_REQUEST_LIMITS).All(&searchResults)
	if err != nil {
		revel.WARN.Printf("[Daily DID] Got error from mongo : [%v].\r\n", err)
		return c.RenderJson(searchResults)
	}
	revel.TRACE.Printf("[Daily DID] send to the client response of %d records.\r\n", len(searchResults))
	return c.RenderJson(searchResults)
}
