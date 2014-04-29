package tests

import (
	"github.com/robfig/revel"
	"net/http"
	"time"
)

var (
	testUser     = "testUser"
	testPassword = "testpassword"
	testAdmin    = "true"
	testDate     = "2014-04-17T01:00:00Z"
	testDID      = "1157"
)

type AppTest struct {
	revel.TestSuite
}

func (t AppTest) Before() {
	println("Set up")
	timenow := time.Now()
	testDate = timenow.Format(time.RFC3339)
}

func (t AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContains("Application Is Ready")
}

func (t AppTest) TestCreateDeleteUser() {
	t.Get("/createuser/" + testUser + "/" + testPassword + "/" + testAdmin)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"status\":200")
	//
	t.Get("/deleteuser/" + testUser + "/" + testPassword)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"status\":200")
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
	t.Get("/cdrs/2013-02-01T00:00:00Z/2015-04-22T01:00:00Z")
	t.AssertOk()
	t.AssertContains("dst")
}

func (t AppTest) TestDailyIncommingCalls() {
	t.Get("/daily/incomming")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
}

func (t AppTest) TestDailyIncommingCallsByDay() {
	t.Get("/daily/incomming/" + testDate)
	t.AssertOk()
	t.AssertContains("id")
}

func (t AppTest) TestDailyIncommingCallsByDayAndUser() {
	t.Get("/daily/incomming/" + testDate)
	t.AssertOk()
	t.AssertContains("id")
}

func (t AppTest) TestDailyIncommingCallsByDayAndUserFailed() {
	t.Get("/daily/incomming/" + testDate + "/81")
	t.AssertOk()
	t.AssertContains("null")
}

func (t AppTest) TestIncommingDidCallsByDay() {
	t.Get("/daily/didincomming/" + testDate)
	t.AssertOk()
	t.AssertContains("id")
}

func (t AppTest) TestIncommingDidCallsByDayAndDID() {
	t.Get("/daily/didincomming/" + testDate + "/" + testDID)
	t.AssertOk()
	t.AssertContains("id")
}

func (t AppTest) After() {
	println("Tear down")
}
