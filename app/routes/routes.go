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

func (p tApp) ListUsers(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.ListUsers", args).Url
}

func (p tApp) UpdateUser(
		username string,
		password string,
		firstname string,
		lastname string,
		isadmin string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "firstname", firstname)
	revel.Unbind(args, "lastname", lastname)
	revel.Unbind(args, "isadmin", isadmin)
	return revel.MainRouter.Reverse("App.UpdateUser", args).Url
}

func (p tApp) CreateUser(
		username string,
		password string,
		firstname string,
		lastname string,
		admin bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "firstname", firstname)
	revel.Unbind(args, "lastname", lastname)
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

func (p tApp) GetDids(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetDids", args).Url
}

func (p tApp) CreateDid(
		id string,
		value string,
		comment string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "value", value)
	revel.Unbind(args, "comment", comment)
	return revel.MainRouter.Reverse("App.CreateDid", args).Url
}

func (p tApp) UpdateDid(
		id string,
		value string,
		comment string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "value", value)
	revel.Unbind(args, "comment", comment)
	return revel.MainRouter.Reverse("App.UpdateDid", args).Url
}

func (p tApp) DeleteDid(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.DeleteDid", args).Url
}

func (p tApp) GetPeers(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetPeers", args).Url
}

func (p tApp) CreatePeer(
		id string,
		value string,
		comment string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "value", value)
	revel.Unbind(args, "comment", comment)
	return revel.MainRouter.Reverse("App.CreatePeer", args).Url
}

func (p tApp) UpdatePeer(
		id string,
		value string,
		comment string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "value", value)
	revel.Unbind(args, "comment", comment)
	return revel.MainRouter.Reverse("App.UpdatePeer", args).Url
}

func (p tApp) DeletePeer(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.DeletePeer", args).Url
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

func (p tDaily) PeersDatas(
		day string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	return revel.MainRouter.Reverse("Daily.PeersDatas", args).Url
}

func (p tDaily) PeerDatas(
		day string,
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Daily.PeerDatas", args).Url
}


type tYearly struct {}
var Yearly tYearly


func (p tYearly) IncommingDidCallsForYearDid(
		year int,
		did string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "year", year)
	revel.Unbind(args, "did", did)
	return revel.MainRouter.Reverse("Yearly.IncommingDidCallsForYearDid", args).Url
}

func (p tYearly) IncommingDidCallsForYearByDid(
		year int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "year", year)
	return revel.MainRouter.Reverse("Yearly.IncommingDidCallsForYearByDid", args).Url
}

func (p tYearly) PeersDatas(
		year int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "year", year)
	return revel.MainRouter.Reverse("Yearly.PeersDatas", args).Url
}

func (p tYearly) PeerDatas(
		year int,
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "year", year)
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Yearly.PeerDatas", args).Url
}


type tMonthly struct {}
var Monthly tMonthly


func (p tMonthly) IncommingDidCallsForMonthDid(
		day string,
		did string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	revel.Unbind(args, "did", did)
	return revel.MainRouter.Reverse("Monthly.IncommingDidCallsForMonthDid", args).Url
}

func (p tMonthly) IncommingDidCallsForMonthByDid(
		day string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	return revel.MainRouter.Reverse("Monthly.IncommingDidCallsForMonthByDid", args).Url
}

func (p tMonthly) PeersDatas(
		day string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	return revel.MainRouter.Reverse("Monthly.PeersDatas", args).Url
}

func (p tMonthly) PeerDatas(
		day string,
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Monthly.PeerDatas", args).Url
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


