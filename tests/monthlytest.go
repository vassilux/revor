package tests

import (
	"github.com/robfig/revel"
	"net/http"
)

var (
	monthlyTestDate = "2014-05-16T01:00:00Z"
	monthlyTestDID  = "1157"
	monthlyTestPeer = "6005"
)

type MonthlyTest struct {
	revel.TestSuite
}

func (t MonthlyTest) Before() {

}

func (t MonthlyTest) TestPeersDatas() {
	t.Get("/monthly/peersdatas/" + monthlyTestDate)
	t.AssertOk()
	t.AssertContains("inCalls")
	t.AssertContains("outCalls")
}

func (t MonthlyTest) TestPeerDatas() {
	t.Get("/monthly/peerdatas/" + monthlyTestDate + "/" + monthlyTestPeer)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("inCalls")
	t.AssertContains("outCalls")
}

func (t MonthlyTest) TestIncommingDidCallsByDay() {
	t.Get("/monthly/didincomming/" + monthlyTestDate)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	println(t.ResponseBody)
	t.AssertContains("id")

}

func (t MonthlyTest) TestIncommingDidCallsByDayAndDID() {
	t.Get("/monthly/didincomming/" + monthlyTestDate + "/" + monthlyTestDID)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("id")
}

func (t MonthlyTest) After() {
	println("Tear down")
}
