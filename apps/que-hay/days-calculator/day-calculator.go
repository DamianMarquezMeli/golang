package main

import (
	"time"
)

type Date struct {
	day   int
	month int
	year  int
}

func main() {
	date1 := Date{
		day:   31,
		month: 4,
		year:  1986,
	}
	date2 := Date{
		day:   30,
		month: 2,
		year:  2345,
	}

	DaysCalculator(date1, date2)
}

func DaysCalculator(start, end Date) int {
	startDate := time.Date(start.year, time.Month(start.month), start.day, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(end.year, time.Month(end.month), end.day, 0, 0, 0, 0, time.UTC)
	days := int(endDate.Sub(startDate).Hours() / 24)
	return days
}
