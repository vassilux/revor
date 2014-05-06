package tests

import (
	"github.com/robfig/revel"
	"net/http"
	"time"
)

/*var (
	testUser     = "testUser"
	testPassword = "testpassword"
	testAdmin    = "true"
	testDate     = "2014-04-17T01:00:00Z"
	testDID      = "1157"
)
*/

type CdrTest struct {
	revel.TestSuite
}

func (t CdrTest) Before() {
	println("Set up")
	timenow := time.Now()
	testDate = timenow.Format(time.RFC3339)
}

func (t CdrTest) TestCdrs() {
	t.Get("/cdrs")
	t.AssertOk()
}

func (t CdrTest) TestCdrsDates() {
	t.Get("/cdrs/2013-02-01T00:00:00Z/2015-04-22T01:00:00Z")
	t.AssertOk()
	t.AssertContains("dst")
}

func (t CdrTest) TestCdrsWithDatesParams() {
	t.Get("/cdrs/startdate,$gte,2013-02-01T00:00:00Z&enddate,$lte,2015-04-22T01:00:00Z")
	t.AssertOk()
	t.AssertContains("_id")
}

func (t CdrTest) TestCdrsWithDatesDispositionParams() {
	t.Get("/cdrs/startdate,$gte,2013-02-01T00:00:00Z&enddate,$lte,2015-04-22T01:00:00Z&disposition,,16")
	t.AssertOk()
	t.AssertContains("_id")
	t.AssertContains("\"disposition\": 16,")
}

func (t CdrTest) TestCdrsWithDestinationParams() {
	t.Get("/cdrs/destination,,6005")
	t.AssertOk()
	t.AssertContains("_id")
	t.AssertContains("\"dst\": \"6005\"")
}

func (t CdrTest) TestCdrsWithCallerIdParams() {
	t.Get("/cdrs/callerid,,9002")
	t.AssertOk()
	t.AssertContains("_id")
	t.AssertContains("\"clid_number\": \"9002\"")
}

func (t CdrTest) TestCdrsWithDirectionParams() {
	t.Get("/cdrs/direction,,2")
	t.AssertOk()
	t.AssertContains("_id")
	t.AssertContains("\"inout_status\": 2")
}

func (t CdrTest) TestCdrsWithFailedParams() {
	t.Get("/cdrs/startdate")
	t.AssertStatus(http.StatusBadRequest)
}

func (t CdrTest) After() {
	println("Tear down")
}
