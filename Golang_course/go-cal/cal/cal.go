package cal

import (
	"fmt"
	"time"
)

// PrintMonth prints a single month calendar
func PrintMonth(year int, month time.Month, showWeekNumbers bool) {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("     %s %d\n", month, year)
	fmt.Println("Su Mo Tu We Th Fr Sa")

	weekday := int(t.Weekday())
	for i := 0; i < weekday; i++ {
		fmt.Print("   ")
	}

	days := daysIn(month, year)
	for day := 1; day <= days; day++ {
		if showWeekNumbers && (day == 1 || int(t.Weekday()) == 0) {
			weekNum := isoWeekNum(year, month, day)
			fmt.Printf("%2d ", weekNum)
		} else if showWeekNumbers {
			fmt.Print("   ")
		}
		fmt.Printf("%2d ", day)
		t = t.AddDate(0, 0, 1)
		if int(t.Weekday()) == 0 {
			fmt.Println()
		}
	}
	fmt.Println("tttt")
}

// PrintMultipleMonths prints n months starting from startYear/startMonth
func PrintMultipleMonths(startYear int, startMonth time.Month, n int, showWeekNumbers bool) {
	year := startYear
	month := startMonth
	for i := 0; i < n; i++ {
		PrintMonth(year, month, showWeekNumbers)
		month++
		if month > 12 {
			month = 1
			year++
		}
	}
}

func daysIn(month time.Month, year int) int {
	t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	return t.Day()
}

func isoWeekNum(year int, month time.Month, day int) int {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	_, week := t.ISOWeek()
	return week
}
