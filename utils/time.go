package utils

import "time"

func StartOfDay(t time.Time, timezone string) time.Time {
	location, _ := time.LoadLocation(timezone)
	year, month, day := t.In(location).Date()

	dayStartTime := time.Date(year, month, day, 0, 0, 0, 0, location)

	return dayStartTime
}

func EndOfDay(t time.Time, timezone string) time.Time {
	location, _ := time.LoadLocation(timezone)
	year, month, day := t.In(location).Date()

	dayEndTime := time.Date(year, month, day, 23, 59, 59, 0, location)

	return dayEndTime
}
