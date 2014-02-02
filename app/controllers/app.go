package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"log"
	"revor/app/modules/mongo"
	"time"
)

type App struct {
	*revel.Controller
	mongo.Mongo
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Cdrs() revel.Result {
	//var ids []int = c.Params.Bind("ids[]", reflect.TypeOf([]int{}))
	//var startParameter time.Time
	//var p = c.Params.Bind("name", reflect.TypeOf(string))
	//
	cdrs := c.MongoDatabase.C("cdrs")
	var startDate = time.Date(2013, 7, 1, 1, 0, 0, 0, time.UTC)
	var endDate = time.Date(2013, 8, 1, 1, 0, 0, 0, time.UTC)
	var searchResults []mongo.Cdr
	//var query = bson.M{"calldate": { $gte: start, $lte: end }}
	//{ $gte: start, $lte: end }
	var limit = 10000
	var query = bson.M{"calldate": bson.M{"$gte": startDate, "$lte": endDate}}
	var err = cdrs.Find(query).Limit(limit).All(&searchResults)
	if err != nil {
		log.Printf("[mongo] Monthly insert failed with error : [%v].", err)
		return c.RenderJson(searchResults)
	}
	return c.RenderJson(searchResults)

}
