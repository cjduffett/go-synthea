package utils

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

// IsValidUnitOfAge returns true if the given unit is a valid unit of age.
// Currently, units of time and age are the same.
func IsValidUnitOfAge(unit string) bool {
	return IsValidUnitOfTime(unit)
}
