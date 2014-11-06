// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	controllers1 "github.com/revel/revel/modules/static/app/controllers"
	_ "github.com/revel/revel/modules/testrunner/app"
	controllers0 "github.com/revel/revel/modules/testrunner/app/controllers"
	_ "revor/app"
	_ "revor/app/broker"
	controllers "revor/app/controllers"
	_ "revor/app/modules/mongo"
	_ "revor/app/modules/utils"
	tests "revor/tests"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					60: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Login",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Logout",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CurrentUser",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ListUsers",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "firstname", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "lastname", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "isadmin", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CreateUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "firstname", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "lastname", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "admin", Type: reflect.TypeOf((*bool)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "username", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "password", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetDids",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CreateDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "value", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "comment", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdateDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "value", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "comment", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeleteDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					253: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "GetPeers",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CreatePeer",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "value", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "comment", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UpdatePeer",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "value", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "comment", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DeletePeer",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					308: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "EventStream",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "EventPoll",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					46: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					69: []string{ 
						"error",
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Cdrs)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Cdrs",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CdrsWithDate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "start", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "end", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CdrDetails",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "uniqueid", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "CdrWithParams",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "params", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					113: []string{ 
					},
					122: []string{ 
					},
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Monthly)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "IncommingDidCallsForMonthDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "did", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingDidCallsForMonthByDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingDidCallsForMonthByMonthDays",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingDidCallsForMonthByMonthDaysAndDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "did", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PeersDatas",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PeerDatas",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingPeerDatasForMonthByMonthDays",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingPeerDatasForMonthByMonthDaysAndPeer",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "peer", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "OutgoingPeerDatasForMonthByMonthDays",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "OutgoingPeerDatasForMonthByMonthDaysAndPeer",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "peer", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Yearly)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "IncommingDidCallsForYearDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "did", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingDidCallsForYearByDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PeersDatas",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PeerDatas",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingDidCallsByMonth",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingDidCallsByMonthAndDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "did", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DidCallsDataByMonthForYear",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DidCallsDataByMonthForYearAndDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "did", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PeersInCallsByMonth",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PeersInCallsByMonthAndPeer",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "peer", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PeersOutCallsByMonth",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PeersOutCallsByMonthAndPeer",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "year", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "peer", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.Daily)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "IncommingCalls",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingCallsByDay",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingCallsByDayUser",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingDidCallsForDayDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "did", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingDidCallsForDayByDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingDidCallsByHourByDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "IncommingDidCallsByHourByDayAndDid",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "did", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PeersDatas",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "PeerDatas",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "day", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"revor/app/models.(*User).Validate": { 
			29: "user.Username",
			37: "user.Username",
		},
		"revor/app/models.ValidatePassword": { 
			45: "password",
		},
	}
	revel.TestSuites = []interface{}{ 
		(*tests.DailyTest)(nil),
		(*tests.PeerTest)(nil),
		(*tests.UserTest)(nil),
		(*tests.CdrTest)(nil),
		(*tests.AppTest)(nil),
		(*tests.MonthlyTest)(nil),
		(*tests.DidTest)(nil),
		(*tests.YearlyTest)(nil),
	}

	revel.Run(*port)
}
