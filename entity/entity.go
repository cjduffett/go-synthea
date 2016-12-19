package entity

import (
	"time"

	"github.com/cjduffett/synthea/records"
)

// Entity is an entity to simulate in Synthea's world.
// Currently the only Entity we simulate is a Patient.
type Entity struct {
	Patient       Patient
	Record        records.Record
	Attributes    map[string]interface{}
	lastWellVisit time.Time
}

/*
// NextWellnessEncounter returns the next wellness
// encounter that should be processed given the current
// simulation time.
func (e *Entity) NextWellnessEncounter(time time.Time) time.Time {
	age := e.Patient.getAgeAtTime(time)
}

func wellnessEncounterSchedule(age int) int {
	// Based on the age, return a factor that becomes the duration
	// between scheduled wellness encounters.
}
*/
