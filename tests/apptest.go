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

func (t AppTest) After() {
	println("Tear down")
}
