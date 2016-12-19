package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type PatientTestSuite struct {
	suite.Suite
	startTime time.Time
	endTime   time.Time
}

func TestPatientTestSuite(t *testing.T) {
	suite.Run(t, new(PatientTestSuite))
}

func (p *PatientTestSuite) SetupSuite() {
	// Set a fixed end time
	p.endTime = time.Date(2016, time.December, 1, 0, 0, 0, 0, time.UTC)
	p.startTime = p.endTime.AddDate(-100, 0, 0)
}

// Tests that the chosen birthdate is in range
func (p *PatientTestSuite) TestPickBirthdateRange() {
	targetAge := 45
	now := time.Date(2016, time.December, 1, 12, 5, 0, 0, time.UTC)
	expectedEarliest := time.Date(1970, time.December, 2, 12, 5, 0, 0, time.UTC)
	expectedLatest := time.Date(1971, time.December, 1, 12, 5, 0, 0, time.UTC)

	birthdate := pickBirthdate(now, targetAge)
	p.True(birthdate.After(expectedEarliest), "Birthdate is more than one year before target age")
	p.True(birthdate.Before(expectedLatest), "Birthdate is too recent to meet target age")
}

func (p *PatientTestSuite) TestPatientGetAge() {
	patient := NewPatient(p.startTime, p.endTime)
	patient.birthDate = time.Date(1994, time.January, 6, 12, 58, 00, 0, time.UTC)

	// Pick a time in the simulation between those two dates
	simTime := time.Date(2016, time.December, 9, 12, 00, 00, 0, time.UTC)

	age := patient.getAgeAtTime(simTime)
	p.Equal(22, age, "Patient born in 1994 expected age in 2016 is 22")
}

func (p *PatientTestSuite) TestPatientPickGender() {
	genders := []string{"Male", "Female"}
	patient := NewPatient(p.startTime, p.endTime)
	p.True(contains(genders, patient.gender), "Invalid gender")
}

func (p *PatientTestSuite) TestPatientPickRace() {
	races := []string{"White", "Hispanic", "Black", "Asian", "Native", "Other"}
	race := pickRace()
	p.True(contains(races, race), "Invalid race")
}

func (p *PatientTestSuite) TestPatientPickEthnicityGivenRace() {
	eths := []string{"Puerto Rican", "Mexican", "Central American", "South American"}
	race := "Hispanic"
	eth := pickEthnicity(race)
	p.True(contains(eths, eth), "Invalid ethnicity")
}

func (p *PatientTestSuite) TestPatientPickBloodTypeGivenRace() {
	typs := []string{"o_positive", "o_negative", "a_positive", "a_negative", "b_positive", "b_negative", "ab_positive", "ab_negative"}
	race := "White"
	typ := pickBloodType(race)
	p.True(contains(typs, typ), "Invalid blood type")
}

func (p *PatientTestSuite) TestPatientPickCurrentAddress() {
	addr := pickCurrentAddress()
	p.NotEmpty(addr.line[0], "Invalid address line[0]")
	p.NotEmpty(addr.state, "Invalid state")
	p.Equal(2, len(addr.state), "State name is not a standard 2 character abbreviation")
	p.NotEmpty(addr.city, "Invalid city")
	p.NotEmpty(addr.postalCode, "Invalid postal code")
	p.Equal(5, len(addr.postalCode), "Invalid postal code")
}

func contains(choices []string, want string) bool {
	// Tests if a slice of strings contains a certain string
	for _, el := range choices {
		if el == want {
			return true
		}
	}
	return false
}
