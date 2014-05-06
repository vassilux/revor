// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tApp struct {}
var App tApp


func (p tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (p tApp) Login(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.Login", args).Url
}

func (p tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).Url
}

func (p tApp) CurrentUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.CurrentUser", args).Url
}

func (p tApp) CreateUser(
		username string,
		password string,
		admin bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "admin", admin)
	return revel.MainRouter.Reverse("App.CreateUser", args).Url
}

func (p tApp) DeleteUser(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.DeleteUser", args).Url
}

func (p tApp) EventStream(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.EventStream", args).Url
}

func (p tApp) EventPoll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.EventPoll", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (p tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (p tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (p tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tStatic struct {}
var Static tStatic


func (p tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (p tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tDaily struct {}
var Daily tDaily


func (p tDaily) IncommingCalls(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Daily.IncommingCalls", args).Url
}

func (p tDaily) IncommingCallsByDay(
		day string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	return revel.MainRouter.Reverse("Daily.IncommingCallsByDay", args).Url
}

func (p tDaily) IncommingCallsByDayUser(
		day string,
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Daily.IncommingCallsByDayUser", args).Url
}

func (p tDaily) IncommingDidCallsForDayDid(
		day string,
		did string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	revel.Unbind(args, "did", did)
	return revel.MainRouter.Reverse("Daily.IncommingDidCallsForDayDid", args).Url
}

func (p tDaily) IncommingDidCallsForDayByDid(
		day string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	return revel.MainRouter.Reverse("Daily.IncommingDidCallsForDayByDid", args).Url
}


type tCdrs struct {}
var Cdrs tCdrs


func (p tCdrs) Cdrs(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Cdrs.Cdrs", args).Url
}

func (p tCdrs) CdrsWithDate(
		start string,
		end string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "start", start)
	revel.Unbind(args, "end", end)
	return revel.MainRouter.Reverse("Cdrs.CdrsWithDate", args).Url
}

func (p tCdrs) CdrDetails(
		uniqueid string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "uniqueid", uniqueid)
	return revel.MainRouter.Reverse("Cdrs.CdrDetails", args).Url
}

func (p tCdrs) CdrWithParams(
		params string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "params", params)
	return revel.MainRouter.Reverse("Cdrs.CdrWithParams", args).Url
}


