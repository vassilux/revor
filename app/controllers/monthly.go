package controllers

import (
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"revor/app/modules/mongo"
)

type Monthly struct {
	App
}

//did parts
//Fetch all incomming calls for given day and did
//return a json stream with calls on success otherwise http status 500
func (c Monthly) IncommingDidCallsForMonthDid(day string, did string) revel.Result {
	revel.TRACE.Printf("[Monthly DID] Get incomming call by month for the date [%s] and did [%s].\r\n", day, did)
	results := mongo.GetDidCallsByMonthAndDid(day, did, c.MongoDatabase)
	revel.TRACE.Printf("[Monthly DID] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//Fetch all incomming calls by did for given day
//return a json stream with calls on success otherwise http status 500
func (c Monthly) IncommingDidCallsForMonthByDid(day string) revel.Result {
	revel.TRACE.Printf("[Monthly DID] Get month incomming call by did for the given date [%s].\r\n", day)
	results := mongo.GetDidMonthCalls(day, c.MongoDatabase)
	revel.TRACE.Printf("[Monthly DID] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//Fetch all datas for peers
//return a json stream with calls on success otherwise http status 500
func (c Monthly) PeersDatas(day string) revel.Result {
	revel.TRACE.Printf("[Monthly PEERSDATAS] Get incomming call for the given date [%s].\r\n", day)
	results := bson.M{}
	inCalls := mongo.GetMonthPeerInCalls(day, c.MongoDatabase)
	outCalls := mongo.GetMonthPeerOutCalls(day, c.MongoDatabase)
	results["inCalls"] = inCalls
	results["outCalls"] = outCalls
	revel.TRACE.Printf("[Monthly PEERSDATAS] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//Fetch all datas for the given day and given user
//return a json stream with calls on success otherwise http status 500
func (c Monthly) PeerDatas(day string, user string) revel.Result {
	revel.TRACE.Printf("[Monthly PEERDATAS] Get incomming call for the given date [%s] and user %s.\r\n", day, user)
	results := bson.M{}
	inCalls := mongo.GetMonthPeerInCallsForUser(day, user, c.MongoDatabase)
	revel.TRACE.Printf("[Monthly PEERSDATAS] Get outgoing call for the given date [%s].\r\n", day)
	outCalls := mongo.GetMonthPeerOutCallsForUser(day, user, c.MongoDatabase)
	results["inCalls"] = inCalls
	results["outCalls"] = outCalls
	revel.TRACE.Printf("[Monthly PEERDATAS] Send to the client response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}
