package entity

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cjduffett/synthea/utils"
	"github.com/icrowley/fake"
)

// Patient is a patient simulated by Synthea.
// TODO: fingerprint
// TODO: multiple births
type Patient struct {
	gender       string
	firstName    string
	lastName     string
	birthDate    time.Time
	race         string
	ethnicity    string
	bloodType    string
	height       float64 // in cm
	weight       float64 // in kg
	address      Address
	placeOfBirth PlaceOfBirth
}

// NewPatient creates a new Patient object
func NewPatient(startDate, endDate time.Time) *Patient {
	// pick a random age for the patient, between 0 and 100
	targetAge := rand.Intn(100)

	// TODO: Config to append ### to the patient's names

	gender := fake.Gender()
	first, last := pickFullName(gender)
	race := pickRace()
	address := pickCurrentAddress()

	return &Patient{
		gender:       gender,
		firstName:    first,
		lastName:     last,
		birthDate:    pickBirthdate(endDate, targetAge),
		race:         race,
		ethnicity:    pickEthnicity(race),
		bloodType:    pickBloodType(race),
		height:       51.0, // Average height at birth
		weight:       3.5,  // Average weight at birth
		address:      address,
		placeOfBirth: pickPlaceOfBirth(address, targetAge),
	}
}

// Address is a patient's full street address
type Address struct {
	line       []string
	city       string
	state      string
	postalCode string
}

// PlaceOfBirth is a patient's place of birth
type PlaceOfBirth struct {
	city    string
	state   string
	country string
}

func (p *Patient) getAgeAtTime(time time.Time) int {
	if time.Before(p.birthDate) {
		panic("Patient has not been born yet")
	}

	years := time.Year() - p.birthDate.Year()
	if time.YearDay() < p.birthDate.YearDay() {
		years--
	}
	return years
}

func pickFullName(gender string) (first, last string) {
	switch gender {
	case "Male":
		first = fake.MaleFirstName()
		last = fake.MaleLastName()
	case "Female":
		first = fake.FemaleFirstName()
		last = fake.FemaleLastName()
	}
	return
}

func pickBirthdate(endDate time.Time, targetAge int) time.Time {
	earliest := endDate.AddDate(-(targetAge + 1), 0, 1)
	latest := endDate.AddDate(-targetAge, 0, 0)
	totalSeconds := int(latest.Sub(earliest).Seconds())
	randomDuration := time.Duration(rand.Intn(totalSeconds))
	return earliest.Add(randomDuration)
}

func pickRace() string {
	races, ok := (Demographics["Race"]).([]utils.Choice)
	if !ok {
		panic("No race demographics provided")
	}
	race, _ := (utils.WeightedChoice(races).Item).(string)
	return race
}

func pickEthnicity(race string) string {
	ethMap, ok := (Demographics["Ethnicity"]).(map[string][]utils.Choice)
	if !ok {
		panic("No ethnicity demographics provided")
	}
	eth, _ := (utils.WeightedChoice(ethMap[race]).Item).(string)
	return eth
}

func pickBloodType(race string) string {
	bloodMap, ok := (Demographics["BloodType"]).(map[string][]utils.Choice)
	if !ok {
		panic("No blood type demographics provided")
	}
	typ, _ := (utils.WeightedChoice(bloodMap[race]).Item).(string)
	return typ
}

func pickCurrentAddress() Address {
	secondaryAddress := ""
	if rand.Float64() < 0.5 {
		secondaryAddress = fmt.Sprintf("APT %s", string(rand.Intn(1000)))
	}

	return Address{
		line:       []string{fake.StreetAddress(), secondaryAddress},
		city:       fake.City(),
		state:      fake.StateAbbrev(),
		postalCode: fake.Zip()[:5], // only use the first 5 digits
	}
}

func pickPlaceOfBirth(address Address, age int) PlaceOfBirth {
	// Based on CDC census data:
	// https://www.census.gov/prod/2011pubs/acsbr10-07.pdf

	changes := []utils.Choice{
		utils.Choice{
			Weight: 0.235,
			Item:   "same",
		},
		utils.Choice{
			Weight: 0.353,
			Item:   "city",
		},
		utils.Choice{
			Weight: 0.129,
			Item:   "country",
		},
		utils.Choice{
			Weight: 0.283,
			Item:   "state",
		},
	}
	change, _ := utils.WeightedChoice(changes).Item.(string)

	pob := PlaceOfBirth{
		city:    address.city,
		state:   address.state,
		country: "United States",
	}

	switch change {
	case "country":
		// Born in a foreign country.
		pob.city = fake.City()
		pob.state = ""
		pob.country = fake.Country()
	case "state":
		// Born in the U.S. but now lives in a different state.
		pob.city = fake.City()
		pob.state = fake.State()
	case "city":
		// Born in the same state but now lives in a different city.
		pob.city = fake.City()
	}
	return pob
}
