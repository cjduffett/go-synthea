package gmf

import (
	"time"

	"github.com/cjduffett/synthea/entity"
)

// Condition is an interface for all condition classes
// that exposes a test() method to test the condition.
type Condition interface {
	test(entity *entity.Entity, time time.Time) bool
}

// AndCondition is a logical AND condition.
type AndCondition struct {
	conditions []Condition
}

func (a *AndCondition) test(entity *entity.Entity, time time.Time) bool {
	// returns true if all conditions are true, false otherwise
	for _, condition := range a.conditions {
		if !condition.test(entity, time) {
			return false
		}
	}
	return true
}

// OrCondition is a logical OR condition.
type OrCondition struct {
	conditions []Condition
}

func (o *OrCondition) test(entity *entity.Entity, time time.Time) bool {
	// returns true if any of the conditions are true
	for _, condition := range o.conditions {
		if condition.test(entity, time) {
			return true
		}
	}
	return false
}

// AtLeastCondition requires that at least n conditions are true.
// If n = 1 AtLeastCondition is equivalent to an OR condition.
// If n = # of conditions AtLeastCondition is equivalent to an AND condition.
type AtLeastCondition struct {
	minimum    int
	conditions []Condition
}

func (al *AtLeastCondition) test(entity *entity.Entity, time time.Time) bool {
	truths := 0
	for _, condition := range al.conditions {
		if condition.test(entity, time) {
			truths++
		}
	}
	if truths >= al.minimum {
		return true
	}
	return false
}

// AtMostCondition requires that at most n conditions are true.
type AtMostCondition struct {
	maximum    int
	conditions []Condition
}

func (am *AtMostCondition) test(entity *entity.Entity, time time.Time) bool {
	truths := 0
	for _, condition := range am.conditions {
		if condition.test(entity, time) {
			truths++
		}
	}
	if truths <= am.maximum {
		return true
	}
	return false
}

// NotCondition negates the given condition.
type NotCondition struct {
	condition Condition
}

func (n *NotCondition) test(entity *entity.Entity, time time.Time) bool {
	return !n.condition.test(entity, time)
}

// GenderCondition tests if the patient is a given gender.
type GenderCondition struct {
	gender string
}

func (g *GenderCondition) test(entity *entity.Entity, time time.Time) bool {
	// TODO: Test against the patient entity
	if g.gender != "M" && g.gender != "F" {
		panic("'%s' is not a valid gender")
	}
	return true
}

// AgeCondition tests if the patient is a certain age.
type AgeCondition struct {
	operator string
	quantity float64
	unit     string
}

func (a *AgeCondition) test(entity *entity.Entity, time time.Time) bool {
	age := normalizeUnitOfTime(a.quantity, a.unit)
	// TODO: Compare against the patient entity
	return compare(age, 100, a.operator)
}

// SocioStatusCondition tests the socioeconomic status of the patient.
type SocioStatusCondition struct {
	category string
}

func (s *SocioStatusCondition) test(entity *entity.Entity, time time.Time) bool {
	// TODO: Test agains the patient entity
	return s.category == "Middle"
}

// RaceCondition tests if the patient is a given race.
type RaceCondition struct {
	race string
}

func (r *RaceCondition) test(entity *entity.Entity, time time.Time) bool {
	return r.race == "Hispanic"
}

// DateCondition compares the current world time to the specified date.
type DateCondition struct {
	operator string
	year     int
}

func (d *DateCondition) test(entity *entity.Entity, time time.Time) bool {
	// TODO: compare against Synthea world
	return compare(d.year, 1990, d.operator)
}

// AttributeCondition compares the specified value against an Attribute
// on the patient entity.
type AttributeCondition struct {
	attribute string
	operator  string
	value     interface{}
}

func (a *AttributeCondition) test(entity *entity.Entity, time time.Time) bool {
	// TODO: compare against patient entity
	return false
}

// SymptomCondition tests the severity of a patient's symptom.
type SymptomCondition struct {
	symptom  string
	operator string
	value    float64
}

func (s *SymptomCondition) test(entity *entity.Entity, time time.Time) bool {
	// TODO: compare against the patient entity
	return false
}

// ObservationCondition tests if an observation has been performed
// on the patient.
type ObservationCondition struct {
	referencedByAttribute string
	codes                 []Code
	operator              string
	value                 float64
}

func (o *ObservationCondition) test(entity *entity.Entity, time time.Time) bool {
	// TODO: compare against patient entity
	return false
}

// PriorStateCondition tests if a state has already been processed.
type PriorStateCondition struct {
	name string
}

func (p *PriorStateCondition) test(entity *entity.Entity, time time.Time) bool {
	// TODO: compare against patient entity
	return false
}

// ActiveCondition tests if a condition previously diagnosed
// on the patient is still active.
type ActiveCondition struct {
	referencedByAttribute string
	codes                 []Code
}

func (a *ActiveCondition) test(entity *entity.Entity, time time.Time) bool {
	// TODO: compare against patient entity
	return false
}

// ActiveCarePlan tests if a care plan previously prescribed
// to the patient is still active.
type ActiveCarePlan struct {
	referencedByAttribute string
	codes                 []Code
}

func (a *ActiveCarePlan) test(entity *entity.Entity, time time.Time) bool {
	// TODO: compare against patient entity
	return false
}

// ActiveMedication tests if medication previously prescribed
// to the patient is still active.
type ActiveMedication struct {
	referencedByAttribute string
	codes                 []Code
}

func (a *ActiveMedication) test(entity *entity.Entity, time time.Time) bool {
	// TODO: compare against patient entity
	return false
}

// TrueCondition always returns true. This is not intended
// for use in the GMF and is for testing purposes only.
type TrueCondition struct{}

func (t *TrueCondition) test(entity *entity.Entity, time time.Time) bool {
	return true
}

// FalseCondition always returns false. This is not intended
// for use in the GMF and is for testing purposes only.
type FalseCondition struct{}

func (t *FalseCondition) test(entity *entity.Entity, time time.Time) bool {
	return false
}

func compare(lhs, rhs int, operator string) bool {
	switch operator {
	case "==":
		return lhs == rhs
	case "!=":
		return lhs != rhs
	case "<":
		return lhs < rhs
	case ">":
		return lhs > rhs
	case "<=":
		return lhs <= rhs
	case ">=":
		return lhs >= rhs
	case "is nil":
		return false
	case "is not nil":
		return false
	default:
		panic("'%s' is not a valid operator")
	}
}

func normalizeUnitOfTime(quantity float64, unit string) int {
	// for comparison we convert all time units to seconds
	switch unit {
	case "years":
		return int(quantity * 3600 * 24 * 30 * 12)
	case "months":
		return int(quantity * 3600 * 24 * 30)
	case "days":
		return int(quantity * 3600 * 24)
	case "hours":
		return int(quantity * 3600)
	case "minutes":
		return int(quantity * 60)
	case "seconds":
		return int(quantity)
	default:
		panic("'%s' is not a valid unit of time")
	}
}
