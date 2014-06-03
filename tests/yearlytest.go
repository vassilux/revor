package tests

import (
	"github.com/revel/revel"
	"net/http"
)

var (
	yearlyTestDate = "2014"
	yearlyTestDID  = "1157"
	yearlyTestPeer = "6005"
)

type YearlyTest struct {
	revel.TestSuite
}

func (t YearlyTest) Before() {

}

func (t YearlyTest) TestPeersDatas() {
	t.Get("/yearly/peerdatas/" + yearlyTestDate)
	t.AssertOk()
	t.AssertContains("inCalls")
	t.AssertContains("outCalls")
}

func (t YearlyTest) TestPeerDatas() {
	t.Get("/yearly/peerdatas/" + yearlyTestDate + "/" + yearlyTestPeer)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("inCalls")
	t.AssertContains("outCalls")
}

func (t YearlyTest) TestIncommingDidCallsByYear() {
	t.Get("/yearly/didincomming/" + yearlyTestDate)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	println(t.ResponseBody)
	t.AssertContains("id")

}

func (t YearlyTest) TestIncommingDidCallsByYearAndDID() {
	t.Get("/yearly/didincomming/" + yearlyTestDate + "/" + yearlyTestDID)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("id")
}

func (t YearlyTest) After() {
	println("Tear down")
}
