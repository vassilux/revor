package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/revel"
	"log"
	"net/http"
	"revor/app/modules/mongo"
	"revor/app/modules/utils"
)

type App struct {
	*revel.Controller
	mongo.Mongo
}

type Response map[string]interface{}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

type LoginResult struct {
	statusCode int
	username   string
	admin      bool
	token      string
}

func (r LoginResult) Apply(req *revel.Request, resp *revel.Response) {
	//the solution is very simple to make json string and inject to the response
	fmt.Fprint(resp.Out, Response{"username": r.username, "admin": r.admin, "status": r.statusCode, "token": r.token})
}

func (c App) Index() revel.Result {
	return c.Render()
}

//inner worker login function
func (c App) ProcessInnerLogin(username, password string) *LoginResult {
	token := utils.GenerateNewToken(c.Controller)
	res := &LoginResult{
		statusCode: http.StatusOK,
		username:   username,
		admin:      false,
		token:      token,
	}
	if len(username) == 0 {
		revel.WARN.Println("username is empty can't process login requst.")
		res.statusCode = http.StatusFound
		return res
	}
	//
	return res
}

//
func (c App) Login(username, password string) revel.Result {
	res := c.ProcessInnerLogin(username, password)
	revel.TRACE.Println(" test : " + c.Session["token"])
	log.Printf("[App] Login session token %s.", c.Session["token"])
	return res
}

func (c App) Logout() revel.Result {
	delete(c.Session, "token")
	return &LoginResult{
		statusCode: http.StatusOK,
	}
}

func (c App) CurrentUser() revel.Result {
	res := &LoginResult{
		statusCode: http.StatusOK,
		username:   "adminTest",
		admin:      true,
		token:      "",
	}
	return res
}
