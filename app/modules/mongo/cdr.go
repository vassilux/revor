package mongo

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"revor/app/models"
	"strconv"
	"time"
)

//Extract the cdr records by the given parameters
func GetCdrs(paramsMap map[string]models.CdrSearchParam, mongoDb *mgo.Database) []bson.M {
	fmt.Printf("Get start [%s] and end [%s] dates for cdrs.\r\n",
		paramsMap["startdate"].Data, paramsMap["enddate"].Data)
	searchResults := []bson.M{}
	var query = bson.M{}

	startDateParam, startDateParamOk := paramsMap["startdate"]
	endDateParam, endDateParamOk := paramsMap["enddate"]

	if startDateParamOk && endDateParamOk {
		startDate, err := time.Parse(time.RFC3339, startDateParam.Data)
		if err != nil {
			panic(err)
		}
		endDate, err := time.Parse(time.RFC3339, endDateParam.Data)
		if err != nil {
			panic(err)
		}
		var datesParams = bson.M{startDateParam.Condition: startDate, endDateParam.Condition: endDate}
		query["call_date"] = datesParams
	}

	disposition, dispositionOk := paramsMap["disposition"]
	if dispositionOk {
		fmt.Printf("Add disposition parameter [%s].\r\n",
			disposition.Data)
		dispositionValue, err := strconv.Atoi(disposition.Data)
		if err != nil {
			panic(err)
		}
		//can be used for more that on values
		//arr := []int{dispositionValue}
		//var dispositionParam = bson.M{"$in": arr}
		query["disposition"] = dispositionValue //dispositionParam
	}

	duration, durationOk := paramsMap["duration"]
	if durationOk {
		fmt.Printf("Add duration parameter [%s].\r\n",
			duration.Data)
		durationValue, err := strconv.Atoi(duration.Data)
		if err != nil {
			panic(err)
		}

		if len(duration.Condition) == 0 {
			query["duration"] = durationValue
		} else {
			query["duration"] = bson.M{duration.Condition: durationValue}
		}

	}
	destination, destinationOk := paramsMap["destination"]
	if destinationOk {
		fmt.Printf("Add destination parameter [%s].\r\n",
			destination.Data)
		if len(destination.Condition) == 0 {
			query["dst"] = destination.Data
		} else {
			query["dst"] = bson.M{"$regex": destination.Data}
		}

	}
	callerId, callerIdOk := paramsMap["callerid"]
	if callerIdOk {
		fmt.Printf("Add callerId parameter [%s].\r\n",
			callerId.Data)
		if len(callerId.Condition) == 0 {
			query["clid_number"] = callerId.Data
		} else {
			query["clid_number"] = bson.M{callerId.Condition: callerId.Data}
		}

	}

	direction, directionOk := paramsMap["direction"]
	if directionOk {
		fmt.Printf("Add direction parameter [%s].\r\n",
			direction.Data)
		directionValue, err := strconv.Atoi(direction.Data)
		if err != nil {
			panic(err)
		}
		query["inout_status"] = directionValue
	}
	did, didOk := paramsMap["did"]
	if didOk {
		fmt.Printf("Add did parameter [%s].\r\n",
			did.Data)
		query["inout_status"] = 2
		query["did"] = did.Data
	}

	cdrs := mongoDb.C("cdrs")
	//
	var limit = 10000
	//hop hop look for
	err := cdrs.Find(query).Limit(limit).All(&searchResults)
	if err != nil {
		panic(err)
	}
	return searchResults
}
