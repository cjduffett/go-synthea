package gmf

import (
	"math/rand"
	"time"

	"github.com/cjduffett/synthea/entity"
)

// State is an interface to all GMF state types.
type State interface {
	process(entity *entity.Entity, time time.Time) bool
	next(entity *entity.Entity, time time.Time) string
}

// Code is the JSON representation of a code.
type Code struct {
	System  string `json:"system"`
	Code    string `json:"code"`
	Display string `json:"display"`
}

// Exact is the JSON representation of an exact quantity.
type Exact struct {
	Quantity int64  `json:"quantity"`
	Unit     string `json:"unit"`
}

// Converts an Exact quantity to a duration of time. The
// Unit field must be a valid unit of time.
func (e *Exact) convertToDuration() time.Duration {
	if isValidUnitOfTime(e.Unit) {
		return convertTimeToDuration(e.Quantity, e.Unit)
	}
	panic("'unit' is not a valid unit of time")
}

// Range is the JSON representation of a ranged quantity.
type Range struct {
	Low  int64  `json:"low"`
	High int64  `json:"high"`
	Unit string `json:"unit"`
}

// Converts an Exact quantity to a duration of time. The
// Unit field must be a valid unit of time.
func (r *Range) convertToDuration() time.Duration {
	var pick int64
	if r.High >= r.Low {
		pick = rand.Int63n(r.High+1-r.Low) + r.Low
		if isValidUnitOfTime(r.Unit) {
			return convertTimeToDuration(pick, r.Unit)
		}
		panic("'unit' is not a valid unit of time")
	}
	panic("'high' cannot be less than 'low'")
}

// InitialState is the initial state of each module. All modules
// should have one and only one Initial state.
type InitialState struct {
	transition Transition
}

func (i *InitialState) process(entity *entity.Entity, time time.Time) bool {
	return true
}

func (i *InitialState) next(entity *entity.Entity, time time.Time) string {
	return i.transition.follow(entity, time)
}

// TerminalState is the last state in a module. Modules may have
// zero or more terminal states. Terminal states do not have
// transitions.
type TerminalState struct {
	transition Transition
}

func (t *TerminalState) process(entity *entity.Entity, time time.Time) bool {
	// By returning false, Terminal blocks the further
	// progression of the module forever, given this entity.
	return false
}

func (t *TerminalState) next(entity *entity.Entity, time time.Time) string {
	panic("The Terminal state does not have a transition.")
}

// SimpleState is the simplest control state. It performs no
// action and makes no changes to the patient's record. It
// just transitions to the next state.
type SimpleState struct {
	name       string
	transition Transition
}

func (s *SimpleState) process(entity *entity.Entity, time time.Time) bool {
	return true
}

func (s *SimpleState) next(entity *entity.Entity, time time.Time) string {
	return s.transition.follow(entity, time)
}

// GuardState blocks module progression until the allow condition is met.
type GuardState struct {
	allow      Condition
	transition Transition
}

func (g *GuardState) process(entity *entity.Entity, time time.Time) bool {
	return g.allow.test(entity, time)
}

func (g *GuardState) next(entity *entity.Entity, time time.Time) string {
	return g.transition.follow(entity, time)
}

// DelayState blocks module progression for a specified length
// of time. If the DelayState is exact, the quantity is stored
// in low.
type DelayState struct {
	startTime  time.Time
	endTime    time.Time
	exact      Exact
	rng        Range
	transition Transition
}

func (d *DelayState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: Delay state processing logic
	return true
}

func (d *DelayState) next(entity *entity.Entity, time time.Time) string {
	return d.transition.follow(entity, time)
}

// EncounterState creates an encounter in the patient's record.
type EncounterState struct {
	wellness   bool
	class      string
	reason     string
	codes      []Code
	transition Transition
}

func (e *EncounterState) process(entity *entity.Entity, time time.Time) bool {

	// TODO: Encounter state processing logic

	// If it's not a wellness encounter, process it immediately:
	// 1. Scan the history for undiagnosed but prior conditions
	// 2. Diagnose them here
	// 3. Transition

	// If it is a wellness encounter:
	// 1. Check when the next scheduled wellness encounter is
	// 2. Delay until that time
	// 3. Based on adherence, either process the encounter or skip it

	return true
}

func (e *EncounterState) next(entity *entity.Entity, time time.Time) string {
	return e.transition.follow(entity, time)
}

// ConditionOnsetState creates a new condition in the patient's record.
type ConditionOnsetState struct {
	targetEncounter   string
	assignToAttribute string
	reason            string
	codes             []Code
	transition        Transition
}

func (c *ConditionOnsetState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: ConditionOnset state logic
	return true
}

func (c *ConditionOnsetState) next(entity *entity.Entity, time time.Time) string {
	return c.transition.follow(entity, time)
}

// ConditionEndState ends the specified condition.
type ConditionEndState struct {
	conditionOnset        string
	referencedByAttribute string
	codes                 []Code
	transition            Transition
}

func (c *ConditionEndState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: ConditionEnd processing logic
	return true
}

func (c *ConditionEndState) next(entity *entity.Entity, time time.Time) string {
	return c.transition.follow(entity, time)
}

// MedicationOrderState order schedules a prescription for the patient.
type MedicationOrderState struct {
	targetEncounter   string
	assignToAttribute string
	reason            string
	codes             []Code
	transition        Transition
}

func (m *MedicationOrderState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: MedicationOrder processing logic
	return true
}

func (m *MedicationOrderState) next(entity *entity.Entity, time time.Time) string {
	return m.transition.follow(entity, time)
}

// MedicationEndState ends a prescription.
type MedicationEndState struct {
	medicationOrder       string
	referencedByAttribute string
	codes                 []Code
	transition            Transition
}

func (m *MedicationEndState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: MedicationEnd processing logic
	return true
}

func (m *MedicationEndState) next(entity *entity.Entity, time time.Time) string {
	return m.transition.follow(entity, time)
}

// CarePlanStartState prescribes a care plan for the patient.
type CarePlanStartState struct {
	targetEncounter   string
	assignToAttribute string
	reason            string
	codes             []Code
	activities        []Code
	transition        Transition
}

func (c *CarePlanStartState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: CarePlanStart processing logic
	return true
}

func (c *CarePlanStartState) next(entity *entity.Entity, time time.Time) string {
	return c.transition.follow(entity, time)
}

// CarePlanEndState ends a prescribed care plan.
type CarePlanEndState struct {
	careplan              string
	referencedByAttribute string
	codes                 []Code
	transition            Transition
}

func (c *CarePlanEndState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: CarePLanEnd processing logic
	return true
}

func (c *CarePlanEndState) next(entity *entity.Entity, time time.Time) string {
	return c.transition.follow(entity, time)
}

// ProcedureState performs a procedure on the patient.
type ProcedureState struct {
	targetEncounter string
	reason          string
	codes           []Code
	transition      Transition
}

func (p *ProcedureState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: Procedure processing logic
	return true
}

func (p *ProcedureState) next(entity *entity.Entity, time time.Time) string {
	return p.transition.follow(entity, time)
}

// ObservationState makes an observation of the patient. If exact,
// the exact quantity is stored in low.
type ObservationState struct {
	targetEncounter string
	exact           Exact
	rng             Range
	unit            string
	codes           []Code
	transition      Transition
}

func (o *ObservationState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: Observation processing logic
	return true
}

func (o *ObservationState) next(entity *entity.Entity, time time.Time) string {
	return o.transition.follow(entity, time)
}

// SymptomState tracks the severity of an arbitrary symptom that
// the patients has. if exact, the exact quantity of the
// symptom is stored in low.
type SymptomState struct {
	symptom    string
	cause      string
	exact      Exact
	rng        Range
	transition Transition
}

func (s *SymptomState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: Symptom processing logic
	return true
}

func (s *SymptomState) next(entity *entity.Entity, time time.Time) string {
	return s.transition.follow(entity, time)
}

// SetAttributeState sets an arbitrary attribute on the patient
// record. Attributes may be strings, states, or numbers.
type SetAttributeState struct {
	attribute  string
	value      interface{}
	transition Transition
}

func (s *SetAttributeState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: SetAttribute processing logic
	return true
}

func (s *SetAttributeState) next(entity *entity.Entity, time time.Time) string {
	return s.transition.follow(entity, time)
}

// Panics if the attribute is not a string
func (s *SetAttributeState) getStringValue() string {
	return s.value.(string)
}

// Panics if the attribute is not numeric
func (s *SetAttributeState) getFloatValue() float64 {
	return s.value.(float64)
}

// Panics if the attribute is not numeric. By default,
// go unmarshal's numerics into float64s. getIntValue
// casts that value as an int.
func (s *SetAttributeState) getIntValue() int {
	return int(s.value.(float64))
}

// Panics if the attribute is not a boolean
func (s *SetAttributeState) getBooleanValue() bool {
	return s.value.(bool)
}

// CounterState increments or decrements the value of an attribute.
// The attribute in question must be a numeric type.
type CounterState struct {
	attribute  string
	action     string
	transition Transition
}

func (c *CounterState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: Counter processing logic
	return true
}

func (c *CounterState) next(entity *entity.Entity, time time.Time) string {
	return c.transition.follow(entity, time)
}

// DeathState results in either an immediate or future death of the
// patient. If exact, the exact quantity is stored in low.
type DeathState struct {
	exact      Exact
	rng        Range
	transition Transition
}

func (d *DeathState) process(entity *entity.Entity, time time.Time) bool {
	// TODO: Death processing logic
	return true
}

func (d *DeathState) next(entity *entity.Entity, time time.Time) string {
	return d.transition.follow(entity, time)
}
