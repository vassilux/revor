package controllers

import (
	"github.com/revel/revel"
	"labix.org/v2/mgo/bson"
	"revor/app/modules/mongo"
)

type Yearly struct {
	App
}

//did parts
//return a json stream with calls on success otherwise http status 500
func (c Yearly) IncommingDidCallsForYearDid(year int, did string) revel.Result {
	revel.TRACE.Printf("[Yearly DID] Get incomming call by month for the date [%s] and did [%s].\r\n", year, did)
	results := mongo.GetDidCallsByYearAndDid(year, did, c.MongoDatabase)
	revel.TRACE.Printf("[Yearly DID] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//return a json stream with calls on success otherwise http status 500
func (c Yearly) IncommingDidCallsForYearByDid(year int) revel.Result {
	revel.TRACE.Printf("[Yearly DID] Get month incomming call by did for the given date [%s].\r\n", year)
	results := mongo.GetDidYearCalls(year, c.MongoDatabase)
	revel.TRACE.Printf("[Yearly DID] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//Fetch all datas for peers
//return a json stream with calls on success otherwise http status 500
func (c Yearly) PeersDatas(year int) revel.Result {
	revel.TRACE.Printf("[Yearly PEERSDATAS] Get incomming call for the given date [%s].\r\n", year)
	results := bson.M{}
	inCalls := mongo.GetYearPeerInCalls(year, c.MongoDatabase)
	outCalls := mongo.GetYearPeerOutCalls(year, c.MongoDatabase)
	results["inCalls"] = inCalls
	results["outCalls"] = outCalls
	revel.TRACE.Printf("[Yearly PEERSDATAS] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//Fetch all datas for the given day and given user
//return a json stream with calls on success otherwise http status 500
func (c Yearly) PeerDatas(year int, user string) revel.Result {
	revel.TRACE.Printf("[Yearly PEERDATAS] Get incomming call for the given date [%s] and user %s.\r\n", year, user)
	results := bson.M{}
	inCalls := mongo.GetYearPeerInCallsForUser(year, user, c.MongoDatabase)
	revel.TRACE.Printf("[Yearly PEERSDATAS] Get outgoing call for the given date [%s].\r\n", year)
	outCalls := mongo.GetYearPeerOutCallsForUser(year, user, c.MongoDatabase)
	results["inCalls"] = inCalls
	results["outCalls"] = outCalls
	revel.TRACE.Printf("[Yearly PEERDATAS] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}
