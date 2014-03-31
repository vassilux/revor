package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"revor/app/models"
	"time"
)

type Cdrs struct {
	App
}

func (c Cdrs) Cdrs() revel.Result {
	cdrs := c.MongoDatabase.C("cdrs")
	var startDate = time.Date(2013, 6, 1, 1, 0, 0, 0, time.UTC)
	var endDate = time.Date(2013, 8, 1, 1, 0, 0, 0, time.UTC)
	var searchResults []models.Cdr
	var limit = 10000
	var query = bson.M{"calldate": bson.M{"$gte": startDate, "$lte": endDate}}
	var err = cdrs.Find(query).Limit(limit).All(&searchResults)
	if err != nil {
		revel.ERROR.Printf("[Cdrs] Got error from mongo : [%v].", err)
		panic(err)
	}
	revel.TRACE.Printf("[Cdrs] send to the client response of %d records.\r\n", len(searchResults))
	return c.RenderJson(searchResults)

}

//Find all calls between two dates
//Dates must be provided into time.RFC 3339 format
//If the provided dates is not valids a panic raised with 500 status code send as the response
func (c Cdrs) CdrsWithDate(start string, end string) revel.Result {
	revel.TRACE.Printf("Get start [%s] and end [%s] dates for cdrs.\r\n", start, end)
	var searchResults []models.Cdr
	cdrs := c.MongoDatabase.C("cdrs")
	startDate, err := time.Parse(time.RFC3339, start)
	if err != nil {
		revel.ERROR.Printf("Failed to parse the start date : [%v].\r\n", err)
		panic(err)
	}
	endDate, err := time.Parse(time.RFC3339, end)
	if err != nil {
		revel.ERROR.Printf("Failed to parse the end date : [%v].\r\n", err)
		//return c.RenderJson(searchResults)
		panic(err)
	}
	//
	var limit = 10000
	var query = bson.M{"calldate": bson.M{"$gte": startDate, "$lte": endDate}}
	err = cdrs.Find(query).Limit(limit).All(&searchResults)
	if err != nil {
		revel.ERROR.Printf("[Cdrs] Monthly insert failed with error : [%v].\r\n", err)
		panic(err)
		//return c.RenderJson(searchResults)
	}
	revel.TRACE.Printf("[Cdrs] send to the client response of %d records.\r\n", len(searchResults))
	return c.RenderJson(searchResults)

}
