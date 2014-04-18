package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/revel"
	"labix.org/v2/mgo/bson"
	"net/http"
	"revor/app/models"
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

type HttpRequestResult struct {
	statusCode int
}

func (r LoginResult) Apply(req *revel.Request, resp *revel.Response) {
	//the solution is very simple to make json string and inject to the response
	fmt.Fprint(resp.Out, Response{"username": r.username, "admin": r.admin, "status": r.statusCode, "token": r.token})
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
			username:   username,
			admin:      false,
			token:      "",
		}
		return res
	}
	//the token for authentificated user
	token := utils.GenerateNewToken(c.Controller)
	//the user authentificated
	res := &LoginResult{
		statusCode: http.StatusOK,
		username:   user.Username,
		admin:      user.IsAdmin,
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
	revel.TRACE.Printf(" [App] loggin result for username [%s] and password [%s] with token [%s]",
		username, password, c.Session["token"])
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

func (c App) CreateUser(username, password string, admin bool) revel.Result {
	user, err := c.getUser(username, password)
	if user != nil {
		revel.ERROR.Printf(" [App CreateUser ] Can not create user [%s] cause user exist get error [%v] and user [%v]",
			username, err, user)
		c.Response.Status = http.StatusConflict
		return c.Render()
	}
	var collection = c.MongoDatabase.C("users")
	//
	newUser := models.User{}
	newUser.Username = username
	newUser.Password = password
	err = collection.Insert(newUser)
	if err != nil {
		revel.ERROR.Printf(" [App CreateUser ] Can not create user [%s], get error [%v]",
			username, err)
		c.Response.Status = http.StatusExpectationFailed
		return c.Render()
	}
	c.Response.Status = http.StatusOK
	//
	return &HttpRequestResult{
		statusCode: http.StatusOK,
	}
}

func (c App) DeleteUser(username, password string) revel.Result {
	user, err := c.getUser(username, password)
	if user == nil {
		revel.ERROR.Printf(" [App DeleteUser ] Can not find user [%s]. Get an error  [%v]",
			username, err)
		c.Response.Status = http.StatusConflict
		return c.Render()
	}
	var collection = c.MongoDatabase.C("users")
	//
	newUser := models.User{}
	newUser.Username = username
	newUser.Password = password
	var selector = bson.M{"username": username, "password": password}
	err = collection.Remove(selector)
	if err != nil {
		revel.ERROR.Printf(" [App DeleteUser] Can not remove user [%s], get error [%v]",
			username, err)
		c.Response.Status = http.StatusExpectationFailed
		return c.Render()
	}
	c.Response.Status = http.StatusOK
	//
	return &HttpRequestResult{
		statusCode: http.StatusOK,
	}
}
