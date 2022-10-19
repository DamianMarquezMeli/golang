package main

import (
	"testing"
	//"time"
)

func TestDaysCalculator(t *testing.T) {

	startDate := Date{
		day:   1,
		month: 1,
		year:  2020,
	}

	endDate := Date{
		day:   10,
		month: 1,
		year:  2020,
	}

	//startDate := time.Date(2020, time.Month(5), 1, 0, 0, 0, 0, time.UTC)
	//endDate := time.Date(2020, time.Month(5), 30, 0, 0, 0, 0, time.UTC)

	got := DaysCalculator(startDate, endDate)
	want := 9

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
