package utils

import "time"

const (
	seconds = "seconds"
	minutes = "minutes"
	hours   = "hours"
	days    = "days"
	weeks   = "weeks"
	years   = "years"
)

var unitsOfTime = map[string]bool{seconds: true, minutes: true, hours: true, days: true, weeks: true, years: true}

// IsValidUnitOfTime returns true if the given unit is a valid unit of time.
func IsValidUnitOfTime(unit string) bool {
	_, found := unitsOfTime[unit]
	return found
}

// ConvertTimeToDuration converts a valid unit of time to a time.Duration.
func ConvertTimeToDuration(unit string) time.Duration {
	switch unit {
	case seconds:
		return time.Second
	case minutes:
		return time.Minute
	case hours:
		return time.Hour
	case days:
		return time.Hour * 24
	case weeks:
		return time.Hour * 24 * 7
	case years:
		return time.Hour * 24 * 365
	default:
		panic("Invalid unit of time")
	}
}
