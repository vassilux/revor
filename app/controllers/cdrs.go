package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"log"
	"revor/app/modules/mongo"
	"time"
)

type Cdrs struct {
	App
}

func (c Cdrs) Cdrs() revel.Result {
	cdrs := c.MongoDatabase.C("cdrs")
	var startDate = time.Date(2013, 6, 1, 1, 0, 0, 0, time.UTC)
	var endDate = time.Date(2013, 8, 1, 1, 0, 0, 0, time.UTC)
	var searchResults []mongo.Cdr
	var limit = 10000
	var query = bson.M{"calldate": bson.M{"$gte": startDate, "$lte": endDate}}
	var err = cdrs.Find(query).Limit(limit).All(&searchResults)
	if err != nil {
		log.Printf("[Cdrs] Got error from mongo : [%v].", err)
		return c.RenderJson(searchResults)
	}
	return c.RenderJson(searchResults)

}

//Find all calls between two dates
//Dates must be provided into time.RFC 3339 format
func (c Cdrs) CdrsWithDate(start string, end string) revel.Result {
	var searchResults []mongo.Cdr
	cdrs := c.MongoDatabase.C("cdrs")
	startDate, err := time.Parse(time.RFC3339, start)
	if err != nil {
		log.Printf("Failed to parse the start date : [%v].", err)
		return c.RenderJson(searchResults)
	}
	endDate, err := time.Parse(time.RFC3339, end)
	if err != nil {
		log.Printf("Failed to parse the end date : [%v].", err)
		return c.RenderJson(searchResults)
	}
	//
	var limit = 10000
	var query = bson.M{"calldate": bson.M{"$gte": startDate, "$lte": endDate}}
	err = cdrs.Find(query).Limit(limit).All(&searchResults)
	if err != nil {
		log.Printf("[Cdrs] Monthly insert failed with error : [%v].", err)
		return c.RenderJson(searchResults)
	}
	log.Printf("[Cdrs] send to the client response of %d records.", len(searchResults))
	return c.RenderJson(searchResults)

}
