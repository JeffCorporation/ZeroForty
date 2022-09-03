package pkg

import (
	"time"
)

const maxNsec = 999999999

//StartOfWeek monday 00:00:00
func StartOfWeek(t time.Time) time.Time {
	weekday := time.Duration(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := t.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return currentZeroDay.Add(-1 * weekday * 24 * time.Hour)
}

//EndOfWeek saturday 23:59:59
func EndOfWeek(t time.Time) time.Time {
	weekday := time.Duration(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := t.Date()
	currentLastDay := time.Date(year, month, day, 23, 59, 59, maxNsec, t.Location())
	return currentLastDay.Add((6 - weekday) * 24 * time.Hour)
}

//func CurrentWeek() (time.Time, time.Time) {
//	from := GetStartDayOfWeek(time.Now())
//
//	return from, from.Add(1 * 7*24*time.Hour)
//}

//WeekRange start and end date of week for a specified date. With a number of previous week
func WeekRange(forTime time.Time, weekCount int) (start time.Time, end time.Time) {
	//Week count must be negative
	if weekCount > 0 {
		weekCount *= -1
	}

	start = StartOfWeek(forTime)
	end = EndOfWeek(forTime)
	end.AddDate(0, 0, weekCount*7)

	return //start, end
}
