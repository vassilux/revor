package tests

import (
	"github.com/robfig/revel"
	"net/http"
)

type AppTest struct {
	revel.TestSuite
}

func (t AppTest) Before() {
	println("Set up")
}

func (t AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContains("Application Is Ready")
}

func (t AppTest) TestLoginOK() {
	t.Get("/login?username=admin&password=crackme")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"status\":200")
}

func (t AppTest) TestLoginFailed() {
	t.Get("/login?username=admin&password=admin")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"status\":403")
}

func (t AppTest) TestCdrs() {
	t.Get("/cdrs")
	t.AssertOk()
}

func (t AppTest) TestCdrsDates() {
	t.Get("/cdrs/2013-02-01T00:00:00Z/2013-11-31T23:59:59Z")
	t.AssertOk()
	t.AssertContains("Dst")
}

func (t AppTest) TestDailyIncommingCalls() {
	t.Get("/daily/incomming")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
}

func (t AppTest) TestDailyIncommingCallsByDay() {
	t.Get("/daily/incomming/2013-11-22T01:00:00Z")
	t.AssertOk()
	t.AssertContains("Id")
}

func (t AppTest) TestDailyIncommingCallsByDayCaller() {
	t.Get("/daily/incomming/2013-11-22T01:00:00Z/139946532")
	t.AssertOk()
	t.AssertContains("Id")
}

func (t AppTest) TestDailyIncommingCallsByDayCallerFailed() {
	t.Get("/daily/incomming/2013-11-22T01:00:00Z/81")
	t.AssertOk()
	t.AssertContains("null")
}

func (t AppTest) After() {
	println("Tear down")
}
