package mongo

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"revor/app/models"
	//"time"
)

//Extract the cdr records by the given parameters
func GetCdrs(paramsMap map[string]models.CdrSearchParam, mongoDb *mgo.Database) []bson.M {
	fmt.Printf("Get start [%s] and end [%s] dates for cdrs.\r\n",
		paramsMap["startdate"].Data, paramsMap["enddate"].Data)
	searchResults := []bson.M{}
	/*searchResults := []bson.M{} //[]models.Cdr
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
	var query = bson.M{"call_date": bson.M{"$gte": startDate, "$lte": endDate}}
	err = cdrs.Find(query).Limit(limit).All(&searchResults)
	if err != nil {
		revel.ERROR.Printf("[Cdrs] Monthly insert failed with error : [%v].\r\n", err)
		panic(err)
		//return c.RenderJson(searchResults)
	}
	revel.TRACE.Printf("[Cdrs] send to the client response of %d records.\r\n", len(searchResults))*/
	return searchResults
}
