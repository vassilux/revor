package tests

import (
	"github.com/robfig/revel"
	"net/http"
)

var (
	dailyTestDate = "2014-05-16T01:00:00Z"
	dailyTestDID  = "1157"
	dailyTestPeer = "6005"
)

type DailyTest struct {
	revel.TestSuite
}

func (t DailyTest) Before() {

}

func (t DailyTest) TestDailyIncommingCalls() {
	t.Get("/daily/incomming")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
}

func (t DailyTest) TestPeersDatas() {
	t.Get("/daily/peersdatas/" + dailyTestDate)
	t.AssertOk()
	t.AssertContains("inCalls")
	t.AssertContains("outCalls")
}

func (t DailyTest) TestPeerDatas() {
	t.Get("/daily/peerdatas/" + dailyTestDate + "/" + dailyTestPeer)
	t.AssertOk()
	t.AssertContains("inCalls")
	t.AssertContains("outCalls")
}

func (t DailyTest) TestDailyIncommingCallsByDay() {
	t.Get("/daily/incomming/" + dailyTestDate)
	t.AssertOk()
	t.AssertContains("id")
}

func (t DailyTest) TestDailyIncommingCallsByDayAndUser() {
	t.Get("/daily/incomming/" + dailyTestDate)
	t.AssertOk()
	t.AssertContains("id")
}

func (t DailyTest) TestDailyIncommingCallsByDayAndUserFailed() {
	t.Get("/daily/incomming/" + dailyTestDate + "/" + dailyTestPeer + "000")
	t.AssertOk()
	t.AssertContains("[]")
}

func (t DailyTest) TestIncommingDidCallsByDay() {
	t.Get("/daily/didincomming/" + dailyTestDate)
	t.AssertOk()
	println(t.ResponseBody)
	t.AssertContains("id")

}

func (t DailyTest) TestIncommingDidCallsByDayAndDID() {
	t.Get("/daily/didincomming/" + dailyTestDate + "/" + dailyTestDID)
	t.AssertOk()
	t.AssertContains("id")
}

func (t DailyTest) After() {
	println("Tear down")
}
