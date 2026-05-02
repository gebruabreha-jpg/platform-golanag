package main

import (
	"flag"
	"time"

	"go-cal/cal"
)

var (
	oneMonth       = flag.Bool("1", false, "show only one month")
	threeMonths    = flag.Bool("3", false, "show previous, current, next month")
	numMonths      = flag.Int("months", 1, "number of months to display")
	showWeekNumber = flag.Bool("week-numbering", false, "show week numbers")
)

func main() {
	flag.Parse()
	now := time.Now()
	year, month := now.Year(), now.Month()

	switch {
	case *oneMonth:
		cal.PrintMonth(year, month, *showWeekNumber)
	case *threeMonths:
		prevMonth := month - 1
		prevYear := year
		if prevMonth < 1 {
			prevMonth = 12
			prevYear--
		}
		nextMonth := month + 1
		nextYear := year
		if nextMonth > 12 {
			nextMonth = 1
			nextYear++
		}
		cal.PrintMonth(prevYear, prevMonth, *showWeekNumber)
		cal.PrintMonth(year, month, *showWeekNumber)
		cal.PrintMonth(nextYear, nextMonth, *showWeekNumber)
	default:
		cal.PrintMultipleMonths(year, month, *numMonths, *showWeekNumber)
	}
}
