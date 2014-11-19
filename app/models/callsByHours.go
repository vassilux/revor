package models

import (
	"fmt"
	"github.com/oleiade/reflections"
	"time"
)

type CallsByHours struct {
	H0  int `json:"0" bson:"0"`
	H1  int `json:"1" bson:"1"`
	H2  int `json:"2" bson:"2"`
	H3  int `json:"3" bson:"3"`
	H4  int `json:"4" bson:"4"`
	H5  int `json:"5" bson:"5"`
	H6  int `json:"6" bson:"6"`
	H7  int `json:"7" bson:"7"`
	H8  int `json:"8" bson:"8"`
	H9  int `json:"9" bson:"9"`
	H10 int `json:"10" bson:"10"`
	H11 int `json:"11" bson:"11"`
	H12 int `json:"12" bson:"12"`
	H13 int `json:"13" bson:"13"`
	H14 int `json:"14" bson:"14"`
	H15 int `json:"15" bson:"15"`
	H16 int `json:"16" bson:"16"`
	H17 int `json:"17" bson:"17"`
	H18 int `json:"18" bson:"18"`
	H19 int `json:"19" bson:"19"`
	H20 int `json:"20" bson:"20"`
	H21 int `json:"21" bson:"21"`
	H22 int `json:"22" bson:"22"`
	H23 int `json:"23" bson:"23"`
}

type HourlyCalls struct {
	Calls CallsByHours `json:"hourlyCalls" bson:"hourlyCalls"`
}
type DayOfWeekCalls struct {
	AnswerWaitTime       int           `json:"answerWaitTime" bson:"answerWaitTime"`
	Calls                int           `json:"calls" bson:"calls"`
	Date                 time.Time     `json:"date" bson:"date"`
	DayOfWeek            int           `json:"dayOfWeek" bson:"dayOfWeek"`
	Disposition          int           `json:"disposition" bson:"disposition"`
	CallsPerHours        []HourlyCalls `json:"callsPerHours" bson:"callsPerHours"`
	SummaryCallsPerHours CallsByHours
}

//
func (D *DayOfWeekCalls) PopulateHoursByDays(results []HourByDaysRecord) {
	for i := 0; i < len(results); i++ {
		hourName := fmt.Sprintf("H%d", i)
		dayName := fmt.Sprintf("Day%d", D.DayOfWeek)
		valueHour, _ := reflections.GetField(D.SummaryCallsPerHours, hourName)
		_ = reflections.SetField(&results[i], dayName, valueHour)
	}
}

func (D *DayOfWeekCalls) MakeSummaryCallsPerHours() {
	for i := 0; i < len(D.CallsPerHours); i++ {
		D.SummaryCallsPerHours.H0 += D.CallsPerHours[i].Calls.H0
		D.SummaryCallsPerHours.H1 += D.CallsPerHours[i].Calls.H1
		D.SummaryCallsPerHours.H2 += D.CallsPerHours[i].Calls.H2
		D.SummaryCallsPerHours.H3 += D.CallsPerHours[i].Calls.H3
		D.SummaryCallsPerHours.H4 += D.CallsPerHours[i].Calls.H4
		D.SummaryCallsPerHours.H5 += D.CallsPerHours[i].Calls.H5
		D.SummaryCallsPerHours.H6 += D.CallsPerHours[i].Calls.H6
		D.SummaryCallsPerHours.H7 += D.CallsPerHours[i].Calls.H7
		D.SummaryCallsPerHours.H8 += D.CallsPerHours[i].Calls.H8
		D.SummaryCallsPerHours.H9 += D.CallsPerHours[i].Calls.H9
		D.SummaryCallsPerHours.H10 += D.CallsPerHours[i].Calls.H10
		D.SummaryCallsPerHours.H11 += D.CallsPerHours[i].Calls.H11
		D.SummaryCallsPerHours.H12 += D.CallsPerHours[i].Calls.H12
		D.SummaryCallsPerHours.H13 += D.CallsPerHours[i].Calls.H13
		D.SummaryCallsPerHours.H14 += D.CallsPerHours[i].Calls.H14
		D.SummaryCallsPerHours.H15 += D.CallsPerHours[i].Calls.H15
		D.SummaryCallsPerHours.H16 += D.CallsPerHours[i].Calls.H16
		D.SummaryCallsPerHours.H17 += D.CallsPerHours[i].Calls.H17
		D.SummaryCallsPerHours.H18 += D.CallsPerHours[i].Calls.H18
		D.SummaryCallsPerHours.H19 += D.CallsPerHours[i].Calls.H19
		D.SummaryCallsPerHours.H20 += D.CallsPerHours[i].Calls.H20
		D.SummaryCallsPerHours.H21 += D.CallsPerHours[i].Calls.H21
		D.SummaryCallsPerHours.H22 += D.CallsPerHours[i].Calls.H22
		D.SummaryCallsPerHours.H23 += D.CallsPerHours[i].Calls.H23
	}
}

type HourByDaysRecord struct {
	Hour string `json:"hour" bson:"hour"`
	Day1 int    `json:"day1" bson:"day1"`
	Day2 int    `json:"day2" bson:"day2"`
	Day3 int    `json:"day3" bson:"day3"`
	Day4 int    `json:"day4" bson:"day4"`
	Day5 int    `json:"day5" bson:"day5"`
	Day6 int    `json:"day6" bson:"day6"`
	Day7 int    `json:"day7" bson:"day7"`
}
