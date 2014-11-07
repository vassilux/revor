package controllers

import (
	"github.com/revel/revel"
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
	results := mongo.GetPeerInCalls(day, c.MongoDatabase)
	revel.TRACE.Printf("[Daily] send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//Fetch all incomming calls for given day and user
//return a json stream with calls on success otherwise http status 500
func (c Daily) IncommingCallsByDayUser(day string, user string) revel.Result {
	revel.TRACE.Printf("[Daily] Get incomming call by day for the date [%s] and caller [%s].\r\n", day, user)
	results := mongo.GetPeerInCallsForUser(day, user, c.MongoDatabase)
	revel.TRACE.Printf("[Daily] send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//did parts
//Fetch all incomming calls for given day and did
//return a json stream with calls on success otherwise http status 500
func (c Daily) IncommingDidCallsForDayDid(day string, did string) revel.Result {
	revel.TRACE.Printf("[Daily DID] Get incomming call by day for the date [%s] and did [%s].\r\n", day, did)
	results := mongo.GetDidCallsByDayAndDid(day, did, c.MongoDatabase)
	revel.TRACE.Printf("[Daily DID] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//Fetch all incomming calls by did for given day
//return a json stream with calls on success otherwise http status 500
func (c Daily) IncommingDidCallsForDayByDid(day string) revel.Result {
	revel.TRACE.Printf("[Daily DID] Get incomming call by did for the given date [%s].\r\n", day)
	results := mongo.GetDidCalls(day, c.MongoDatabase)
	revel.TRACE.Printf("[Daily DID] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//Fetch all incomming calls by did for given day
//return a json stream with calls on success otherwise http status 500
func (c Daily) IncommingDidCallsByHourByDid(day string) revel.Result {
	revel.TRACE.Printf("[Daily DID calls by hour] Get incomming call by hour by did for the given date [%s].\r\n", day)
	results := mongo.GetDidCallsByHours(day, c.MongoDatabase)
	revel.TRACE.Printf("[Daily DID Daily DID calls by hour] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//Fetch all incomming calls by did for given day
//return a json stream with calls on success otherwise http status 500
func (c Daily) IncommingDidCallsByHourByDayAndDid(day string, did string) revel.Result {
	revel.TRACE.Printf("[Daily DID calls by hour and Day by hour] for the given date [%s].\r\n", day)
	results := mongo.GetDidCallsByHoursAndDid(day, did, c.MongoDatabase)
	revel.TRACE.Printf("[Daily DID calls by hour and Day by hour] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

func (c Daily) DidGenStatsByDay(day string, tmp int) revel.Result {
	revel.TRACE.Printf("[Daily DID] Get general stats for the given date [%s].\r\n", day)
	results := mongo.GetDidGenStatsByDayAndDid(day, "", c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Daily) DidGenStatsByDayAndDid(day string, did string, tmp int) revel.Result {
	revel.TRACE.Printf("[Daily DID] Get general stats for the given date [%s] and did [%s].\r\n", day, did)
	results := mongo.GetDidGenStatsByDayAndDid(day, did, c.MongoDatabase)
	return c.RenderJson(results)
}

//Fetch all datas for peers
//return a json stream with calls on success otherwise http status 500
func (c Daily) PeersDatas(day string) revel.Result {
	revel.TRACE.Printf("[Daily PEERSDATAS] Get incomming call for the given date [%s].\r\n", day)
	results := bson.M{}
	inCalls := mongo.GetPeerInCalls(day, c.MongoDatabase)
	outCalls := mongo.GetPeerOutCalls(day, c.MongoDatabase)
	hourlyInCalls := mongo.GetPeerInCallsByHours(day, c.MongoDatabase)
	hourlyOutCalls := mongo.GetPeerOutCallsByHours(day, c.MongoDatabase)
	results["inCalls"] = inCalls
	results["outCalls"] = outCalls
	results["hourlyInCalls"] = hourlyInCalls
	results["hourlyOutCalls"] = hourlyOutCalls
	revel.TRACE.Printf("[Daily PEERSDATAS] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//Fetch all datas for the given day and given user
//return a json stream with calls on success otherwise http status 500
func (c Daily) PeerDatas(day string, user string) revel.Result {
	revel.TRACE.Printf("[Daily PEERDATAS] Get incomming call for the given date [%s].\r\n", day)
	results := bson.M{}
	inCalls := mongo.GetPeerInCallsForUser(day, user, c.MongoDatabase)
	revel.TRACE.Printf("[Daily PEERSDATAS] Get outgoing call for the given date [%s].\r\n", day)
	outCalls := mongo.GetPeerOutCallsForUser(day, user, c.MongoDatabase)
	hourlyInCalls := mongo.GetPeerInCallsByHoursAndPeer(day, user, c.MongoDatabase)
	hourlyOutCalls := mongo.GetPeerOutCallsByHoursAndPeer(day, user, c.MongoDatabase)
	results["inCalls"] = inCalls
	results["outCalls"] = outCalls
	results["hourlyInCalls"] = hourlyInCalls
	results["hourlyOutCalls"] = hourlyOutCalls
	revel.TRACE.Printf("[Daily PEERDATAS] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

func (c Daily) IncommingPeerGenStatsByDay(day string, tmp int) revel.Result {
	revel.TRACE.Printf("[Daily Peer] Get general incomming stats for the given date [%s].\r\n", day)
	results := mongo.GetPeerGenStatsByDayAndPeer(day, "", "in", c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Daily) IncommingPeerGenStatsByDayAndPeer(day string, peer string, tmp int) revel.Result {
	revel.TRACE.Printf("[Daily Peer] Get general incomming stats for the given date [%s] and peer [%s].\r\n", day, peer)
	results := mongo.GetPeerGenStatsByDayAndPeer(day, peer, "in", c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Daily) OutgoingPeerGenStatsByDay(day string, tmp int) revel.Result {
	revel.TRACE.Printf("[Daily Peer] Get general outgoing stats for the given date [%s].\r\n", day)
	results := mongo.GetPeerGenStatsByDayAndPeer(day, "", "out", c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Daily) OutgoingPeerGenStatsByDayAndPeer(day string, peer string, tmp int) revel.Result {
	revel.TRACE.Printf("[Daily Peer] Get general outgoing stats for the given date [%s] and peer [%s].\r\n", day, peer)
	results := mongo.GetPeerGenStatsByDayAndPeer(day, peer, "out", c.MongoDatabase)
	return c.RenderJson(results)
}
