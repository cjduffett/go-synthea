package gmf

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
func isValidUnitOfTime(unit string) bool {
	_, found := unitsOfTime[unit]
	return found
}

// ConvertTimeToDuration converts a valid unit of time to a time.Duration.
func convertTimeToDuration(quantity int64, unit string) time.Duration {
	var factor time.Duration
	switch unit {
	case seconds:
		factor = time.Second
	case minutes:
		factor = time.Minute
	case hours:
		factor = time.Hour
	case days:
		factor = time.Hour * 24
	case weeks:
		factor = time.Hour * 24 * 7
	case years:
		factor = time.Hour * 24 * 365
	default:
		panic("Invalid unit of time")
	}
	return factor * time.Duration(quantity)
}
