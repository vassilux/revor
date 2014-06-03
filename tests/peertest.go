package tests

import (
	"github.com/revel/revel"
	"net/http"
	"net/url"
)

type PeerTest struct {
	revel.TestSuite
}

func (t PeerTest) Before() {
	println("Set up")
}

func (t PeerTest) After() {
	println("Tear down")
}

func (t PeerTest) TestListPeers() {
	t.Get("/peers/")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"id\":")
	t.AssertContains("comment")

}

func (t PeerTest) TestCRUDPeer() {
	datas := url.Values{}
	datas.Add("id", "8181")
	datas.Add("value", "8181")
	datas.Add("comment", "my comment")
	t.PostForm("/createpeer", datas)
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"status\":200")
	//
	datasUpdate := url.Values{}
	datasUpdate.Add("id", "8181")
	datasUpdate.Add("value", "8181 | AFFA")
	datasUpdate.Add("comment", "udpate comment")
	t.PostForm("/updatepeer", datasUpdate)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"status\":200")
	//
	t.Get("/deletepeer/8181")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"status\":200")
}
