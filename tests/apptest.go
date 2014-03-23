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
}

func (t AppTest) TestLoginFailed() {
	t.Get("/login?username=admin&password=admin")
	t.AssertOk()
	t.AssertContains("\"status\":403")
}

func (t AppTest) After() {
	println("Tear down")
}
