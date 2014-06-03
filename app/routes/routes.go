// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Login(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.Login", args).Url
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).Url
}

func (_ tApp) CurrentUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.CurrentUser", args).Url
}

func (_ tApp) ListUsers(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.ListUsers", args).Url
}

func (_ tApp) UpdateUser(
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

func (_ tApp) CreateUser(
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

func (_ tApp) DeleteUser(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.DeleteUser", args).Url
}

func (_ tApp) GetDids(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetDids", args).Url
}

func (_ tApp) CreateDid(
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

func (_ tApp) UpdateDid(
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

func (_ tApp) DeleteDid(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.DeleteDid", args).Url
}

func (_ tApp) GetPeers(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetPeers", args).Url
}

func (_ tApp) CreatePeer(
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

func (_ tApp) UpdatePeer(
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

func (_ tApp) DeletePeer(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.DeletePeer", args).Url
}

func (_ tApp) EventStream(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.EventStream", args).Url
}

func (_ tApp) EventPoll(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.EventPoll", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
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


func (_ tDaily) IncommingCalls(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Daily.IncommingCalls", args).Url
}

func (_ tDaily) IncommingCallsByDay(
		day string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	return revel.MainRouter.Reverse("Daily.IncommingCallsByDay", args).Url
}

func (_ tDaily) IncommingCallsByDayUser(
		day string,
		user string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("Daily.IncommingCallsByDayUser", args).Url
}

func (_ tDaily) IncommingDidCallsForDayDid(
		day string,
		did string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	revel.Unbind(args, "did", did)
	return revel.MainRouter.Reverse("Daily.IncommingDidCallsForDayDid", args).Url
}

func (_ tDaily) IncommingDidCallsForDayByDid(
		day string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	return revel.MainRouter.Reverse("Daily.IncommingDidCallsForDayByDid", args).Url
}

func (_ tDaily) PeersDatas(
		day string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	return revel.MainRouter.Reverse("Daily.PeersDatas", args).Url
}

func (_ tDaily) PeerDatas(
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


func (_ tYearly) IncommingDidCallsForYearDid(
		year int,
		did string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "year", year)
	revel.Unbind(args, "did", did)
	return revel.MainRouter.Reverse("Yearly.IncommingDidCallsForYearDid", args).Url
}

func (_ tYearly) IncommingDidCallsForYearByDid(
		year int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "year", year)
	return revel.MainRouter.Reverse("Yearly.IncommingDidCallsForYearByDid", args).Url
}

func (_ tYearly) PeersDatas(
		year int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "year", year)
	return revel.MainRouter.Reverse("Yearly.PeersDatas", args).Url
}

func (_ tYearly) PeerDatas(
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


func (_ tMonthly) IncommingDidCallsForMonthDid(
		day string,
		did string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	revel.Unbind(args, "did", did)
	return revel.MainRouter.Reverse("Monthly.IncommingDidCallsForMonthDid", args).Url
}

func (_ tMonthly) IncommingDidCallsForMonthByDid(
		day string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	return revel.MainRouter.Reverse("Monthly.IncommingDidCallsForMonthByDid", args).Url
}

func (_ tMonthly) PeersDatas(
		day string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "day", day)
	return revel.MainRouter.Reverse("Monthly.PeersDatas", args).Url
}

func (_ tMonthly) PeerDatas(
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


func (_ tCdrs) Cdrs(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Cdrs.Cdrs", args).Url
}

func (_ tCdrs) CdrsWithDate(
		start string,
		end string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "start", start)
	revel.Unbind(args, "end", end)
	return revel.MainRouter.Reverse("Cdrs.CdrsWithDate", args).Url
}

func (_ tCdrs) CdrDetails(
		uniqueid string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "uniqueid", uniqueid)
	return revel.MainRouter.Reverse("Cdrs.CdrDetails", args).Url
}

func (_ tCdrs) CdrWithParams(
		params string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "params", params)
	return revel.MainRouter.Reverse("Cdrs.CdrWithParams", args).Url
}


