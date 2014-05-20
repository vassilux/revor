package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"net/http"
	"revor/app/broker"
	"revor/app/models"
	"revor/app/modules/mongo"
	"revor/app/modules/utils"
	"strconv"
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
	userName   string
	firstName  string
	lastName   string
	isAdmin    bool
	token      string
}

type HttpRequestResult struct {
	statusCode int
}

func (r LoginResult) Apply(req *revel.Request, resp *revel.Response) {
	//the solution is very simple to make json string and inject to the response
	fmt.Fprint(resp.Out, Response{"userName": r.userName, "firstName": r.firstName,
		"lastName": r.lastName, "isAdmin": r.isAdmin,
		"status": r.statusCode, "token": r.token})
}

func (r HttpRequestResult) Apply(req *revel.Request, resp *revel.Response) {
	//the solution is very simple to make json string and inject to the response
	fmt.Fprint(resp.Out, Response{"status": r.statusCode})
}

func (c App) Index() revel.Result {
	return c.Render()
}

//the authentificaiton for given username and password
//return LoginResult object
func (c App) processInnerLogin(username, password string) *LoginResult {
	var user *models.User
	user, err := c.getUser(username, password)

	if err != nil {
		res := &LoginResult{
			statusCode: http.StatusForbidden,
			userName:   username,
			firstName:  "",
			lastName:   "",
			isAdmin:    false,
			token:      "",
		}
		return res
	}
	//the token for authentificated user
	token := utils.GenerateNewToken(c.Controller)
	//the user authentificated
	res := &LoginResult{
		statusCode: http.StatusOK,
		userName:   username,
		firstName:  user.FirstName,
		lastName:   user.LastName,
		isAdmin:    user.IsAdmin,
		token:      token,
	}
	//
	return res
}

func (c App) getUser(username, password string) (*models.User, error) {
	user := models.User{}
	var collection = c.MongoDatabase.C("users")
	var err = collection.Find(bson.M{"username": username, "password": password}).One(&user)
	if err != nil {
		revel.ERROR.Printf(" [App] User not find [%s] with password [%s]",
			username, password)
		return nil, err
	}

	return &user, nil
}

//
func (c App) Login(username, password string) revel.Result {
	res := c.processInnerLogin(username, password)
	c.Session["username"] = res.userName
	c.Session["token"] = res.token
	adminString := strconv.FormatBool(res.isAdmin)
	c.Session["admin"] = adminString

	revel.TRACE.Printf(" [App] loggin result for username [%s] and password [%s]",
		username, password)
	return res
}

func (c App) Logout() revel.Result {
	delete(c.Session, "username")
	delete(c.Session, "token")
	delete(c.Session, "admin")
	return &LoginResult{
		statusCode: http.StatusOK,
	}
}

func (c App) CurrentUser() revel.Result {
	ok := c.Session["token"]
	if len(ok) == 0 {
		res := &LoginResult{
			statusCode: http.StatusForbidden,
		}
		return res
	}
	admin, _ := strconv.ParseBool(c.Session["admin"])
	res := &LoginResult{
		statusCode: http.StatusOK,
		userName:   c.Session["username"],
		isAdmin:    admin,
		token:      c.Session["token"],
	}

	return res
}

func (c App) ListUsers() revel.Result {
	results := mongo.ListUsers(c.MongoDatabase)
	return c.RenderJson(results)
}

func (c App) UpdateUser(username, password, firstname, lastname, isadmin string) revel.Result {
	admin, _ := strconv.ParseBool(isadmin)
	revel.TRACE.Printf(" [App UpdateUser ] Try update user [%s] , [%s], [%s], [%s], [%s]",
		username, password, firstname, lastname, isadmin)
	err := mongo.UpdateUser(username, password, firstname, lastname, admin, c.MongoDatabase)
	if err != nil {
		revel.ERROR.Printf(" [App UpdateUser ] Can't update user [%s] , [%s], [%s], [%s], [%s]",
			username, password, firstname, lastname, isadmin)
		c.Response.Status = http.StatusNotFound
		return c.Render()
	}
	c.Response.Status = http.StatusOK
	return c.Render()
}

func (c App) CreateUser(username, password, firstname, lastname string, admin bool) revel.Result {
	err := mongo.CreateUser(username, password, firstname, lastname, admin, c.MongoDatabase)
	if err != nil {
		revel.ERROR.Printf(" [App CreateUser ] Can not create user [%s] cause user exist get error [%v]",
			username, err)
		c.Response.Status = http.StatusConflict
		return c.Render()
	}
	//
	return &HttpRequestResult{
		statusCode: http.StatusOK,
	}
}

func (c App) DeleteUser(username, password string) revel.Result {
	err := mongo.DeleteUser(username, password, c.MongoDatabase)
	if err != nil {
		revel.ERROR.Printf(" [App DeleteUser ] Can not delete user [%s] cause user exist get error [%v]",
			username, err)
		c.Response.Status = http.StatusConflict
		return c.Render()
	}
	//
	return &HttpRequestResult{
		statusCode: http.StatusOK,
	}
}

func (c App) GetDids() revel.Result {
	revel.TRACE.Printf("[App GetDids].\r\n")
	results := mongo.GetDids(c.MongoDatabase)
	return c.RenderJson(results)
}

func (c App) GetPeers() revel.Result {
	revel.TRACE.Printf("[App GetPeers].\r\n")
	results := mongo.GetPeers(c.MongoDatabase)
	return c.RenderJson(results)
}

func (c App) EventStream() revel.Result {
	w := c.Response.Out
	messageChan := make(chan string)
	broker.MyBroker.NewClients <- messageChan
	defer func() {
		broker.MyBroker.DefunctClients <- messageChan
	}()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	for {
		var data bytes.Buffer
		select {
		case msg := <-messageChan:
			data.WriteString(fmt.Sprintf("event: %s\n", "dashboard_event"))
			data.WriteString(fmt.Sprintf("data: %s\n", msg))
			w.Write(data.Bytes())
			return nil
		}
	}
	return nil
}

func (c App) EventPoll() revel.Result {
	messageChan := make(chan string)
	broker.MyBroker.NewClients <- messageChan
	defer func() {
		broker.MyBroker.DefunctClients <- messageChan
	}()
	//
	for {
		select {
		case msg := <-messageChan:
			return c.RenderJson(Response{"data": msg})
		}
	}
	return nil
}
