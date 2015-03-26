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

//
//return a json stream with calls on success otherwise http status 500
func (c Yearly) IncommingDidCallsByMonth(year int) revel.Result {
	revel.TRACE.Printf("[Yearly DID] Get dids incomming calls groupped by month for the given year [%s].\r\n", year)
	results := mongo.GetDidYearCallsByMonth(year, c.MongoDatabase)
	revel.TRACE.Printf("[Yearly DID] Get dids incomming calls groupped by month response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//return a json stream with calls on success otherwise http status 500
func (c Yearly) IncommingDidCallsByMonthAndDid(year int, did string) revel.Result {
	revel.TRACE.Printf("[Yearly DID] Get dids incomming calls groupped by month for the given year [%s] and did [%s].\r\n", year, did)
	results := mongo.GetDidYearCallsByMonthAndDid(year, did, c.MongoDatabase)
	revel.TRACE.Printf("[Yearly DID] Get dids incomming calls groupped by month response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

func (c Yearly) DidCallsDataByMonthForYear(year int, tm int) revel.Result {
	revel.TRACE.Printf("[Yearly DID] Get calls data for dids for the given year [%d].\r\n", year)
	results := mongo.GetDidCallsDataByMonthForYearAndDid(year, "", c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Yearly) DidCallsDataByMonthForYearAndDid(year int, did string, tm int) revel.Result {
	revel.TRACE.Printf("[Yearly DID] Get calls data for dids for the given year [%d].\r\n", year)
	results := mongo.GetDidCallsDataByMonthForYearAndDid(year, did, c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Yearly) PeersInCallsByMonth(year int) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get peers incomming calls groupped by month for the given year [%d].\r\n", year)
	results := mongo.GetPeerYearInOutCallsByMonthAndPeer(year, "", "in", c.MongoDatabase)
	revel.TRACE.Printf("[Yearly Peer] Get peers incomming calls groupped by month response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

func (c Yearly) PeersInCallsByMonthAndPeer(year int, peer string) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get peers incomming calls groupped by month for the given year [%d] and peer[%s].\r\n", year, peer)
	results := mongo.GetPeerYearInOutCallsByMonthAndPeer(year, peer, "in", c.MongoDatabase)
	revel.TRACE.Printf("[Yearly Peer] Get peers incomming calls groupped by month response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

func (c Yearly) PeersOutCallsByMonth(year int) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get peers outging calls groupped by month for the given year [%d].\r\n", year)
	results := mongo.GetPeerYearInOutCallsByMonthAndPeer(year, "", "out", c.MongoDatabase)
	revel.TRACE.Printf("[Yearly Peer] Get peers outgoing calls groupped by month response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

func (c Yearly) PeersOutCallsByMonthAndPeer(year int, peer string) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get peer incomming calls groupped by month for the given year [%s] and peer[%s].\r\n", year, peer)
	results := mongo.GetPeerYearInOutCallsByMonthAndPeer(year, peer, "out", c.MongoDatabase)
	revel.TRACE.Printf("[Yearly Peer] Get peer incomming calls groupped by month response of %d records.\r\n", len(results))
	return c.RenderJson(results)
}

//
func (c Yearly) PeerInCallsDataByMonthForYear(year int, tm int) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get calls data for dids for the given year [%d].\r\n", year)
	results := mongo.GetPeerInCallsDataByMonthForYearAndPeer(year, "", c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Yearly) PeerInCallsDataByMonthForYearAndDid(year int, did string, tm int) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get calls data for dids for the given year [%d].\r\n", year)
	results := mongo.GetPeerInCallsDataByMonthForYearAndPeer(year, did, c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Yearly) PeerOutCallsDataByMonthForYear(year int, tm int) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get calls data for dids for the given year [%d].\r\n", year)
	results := mongo.GetPeerOutCallsDataByMonthForYearAndPeer(year, "", c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Yearly) PeerOutCallsDataByMonthForYearAndDid(year int, did string, tm int) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get calls data for dids for the given year [%d].\r\n", year)
	results := mongo.GetPeerOutCallsDataByMonthForYearAndPeer(year, did, c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Yearly) PeerGetInDispositionByYear(year int, tmp int) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get incalls disposition by year [%d].\r\n", year)
	results := mongo.GetPeerDispositionByYear(year, "in", "", c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Yearly) PeerGetInDispositionByYearAndPeer(year int, peer string, tmp int) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get incalls disposition by year [%d] and peer [%s].\r\n", year, peer)
	results := mongo.GetPeerDispositionByYear(year, "in", peer, c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Yearly) PeerGetOutDispositionByYear(year int, tmp int) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get outcalls disposition by year [%d].\r\n", year)
	results := mongo.GetPeerDispositionByYear(year, "out", "", c.MongoDatabase)
	return c.RenderJson(results)
}

func (c Yearly) PeerGetOutDispositionByYearAndPeer(year int, peer string, tmp int) revel.Result {
	revel.TRACE.Printf("[Yearly Peer] Get outcalls disposition by year [%d] and peer [%s].\r\n", year, peer)
	results := mongo.GetPeerDispositionByYear(year, "out", peer, c.MongoDatabase)
	return c.RenderJson(results)
}
