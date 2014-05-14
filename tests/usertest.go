package tests

import (
	"github.com/robfig/revel"
	"net/http"
	"net/url"
)

type UserTest struct {
	revel.TestSuite
}

func (t UserTest) Before() {
	println("Set up")
}

func (t UserTest) After() {
	println("Tear down")
}

func (t UserTest) TestListUsers() {
	t.Get("/users/")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	t.AssertContains("\"isAdmin\": true")

}

func (t UserTest) TestCRUDUsers() {
	datas := url.Values{}
	datas.Add("username", "toto")
	datas.Add("password", "crackme")
	datas.Add("firstname", "firstname")
	datas.Add("lastname", "lastname")
	datas.Add("admin", "true")
	t.PostForm("/createuser", datas)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	//
	datasUpdate := url.Values{}
	datasUpdate.Add("username", "toto")
	datasUpdate.Add("password", "crackme1")
	datasUpdate.Add("firstname", "firstname1")
	datasUpdate.Add("lastname", "lastname1")
	datasUpdate.Add("admin", "false")
	t.PostForm("/updateuser", datasUpdate)
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
	//
	t.Get("/deleteuser/toto/crackme1")
	t.AssertOk()
	t.AssertStatus(http.StatusOK)
}
