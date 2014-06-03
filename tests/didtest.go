package tests

import (
	"github.com/revel/revel"
	"net/http"
	"net/url"
)

type DidTest struct {
	revel.TestSuite
}

func (t DidTest) Before() {
	println("Set up")
}

func (t DidTest) After() {
	println("Tear down")
}

func (t DidTest) TestListDids() {
	t.Get("/dids/")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"id\":")
	t.AssertContains("comment")

}

func (t DidTest) TestCRUDDid() {
	datas := url.Values{}
	datas.Add("id", "8181")
	datas.Add("value", "8181")
	datas.Add("comment", "my comment")
	t.PostForm("/createdid", datas)
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"status\":200")
	//
	datasUpdate := url.Values{}
	datasUpdate.Add("id", "8181")
	datasUpdate.Add("value", "8181 | AFFA")
	datasUpdate.Add("comment", "udpate comment")
	t.PostForm("/updatedid", datasUpdate)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"status\":200")
	//
	t.Get("/deletedid/8181")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"status\":200")
}
