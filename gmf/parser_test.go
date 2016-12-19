package gmf

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}

// ============================================================================
// TEST VALID STATES
// ============================================================================
// This section tests parsing states that we know are valid and complete.

func (suite *ParserTestSuite) TestParseStatesIntoCorrectTypes() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	var ok bool

	_, ok = gmf.modules[0].states["Initial"].(*InitialState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Terminal"].(*TerminalState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Simple"].(*SimpleState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Guard"].(*GuardState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Exact_Delay"].(*DelayState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Range_Delay"].(*DelayState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["ConditionOnset"].(*ConditionOnsetState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Encounter"].(*EncounterState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["MedicationOrder"].(*MedicationOrderState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["CarePlanStart"].(*CarePlanStartState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Procedure"].(*ProcedureState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Condition_End_By_Attribute"].(*ConditionEndState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Condition_End_By_Name"].(*ConditionEndState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Condition_End_By_Code"].(*ConditionEndState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Medication_End_By_Attribute"].(*MedicationEndState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Medication_End_By_Name"].(*MedicationEndState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Medication_End_By_Code"].(*MedicationEndState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["CarePlan_End_By_Attribute"].(*CarePlanEndState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["CarePlan_End_By_Name"].(*CarePlanEndState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["CarePlan_End_By_Code"].(*CarePlanEndState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Exact_Observation"].(*ObservationState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Range_Observation"].(*ObservationState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Exact_Symptom"].(*SymptomState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Range_Symptom"].(*SymptomState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["SetAttributeString"].(*SetAttributeState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["SetAttributeNumeric"].(*SetAttributeState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["SetAttributeBoolean"].(*SetAttributeState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["SetAttributeNil"].(*SetAttributeState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Immediate_Death"].(*DeathState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Exact_Death"].(*DeathState)
	suite.True(ok)

	_, ok = gmf.modules[0].states["Range_Death"].(*DeathState)
	suite.True(ok)
}

// Tests the simplest states: Initial, Terminal, and Simple
func (suite *ParserTestSuite) TestParseSimpleStates() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	initial, _ := gmf.modules[0].states["Initial"].(*InitialState)
	suite.Equal(directTransition("Simple"), initial.transition)

	simple, _ := gmf.modules[0].states["Simple"].(*SimpleState)
	suite.Equal(directTransition("Guard"), simple.transition)

	terminal, _ := gmf.modules[0].states["Terminal"].(*TerminalState)
	suite.Equal(&TerminalState{}, terminal)
}

func (suite *ParserTestSuite) TestParseGuardState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	guard, _ := gmf.modules[0].states["Guard"].(*GuardState)
	suite.Equal(directTransition("Exact_Delay"), guard.transition)

	allow := &GenderCondition{gender: "F"}
	suite.Equal(allow, guard.allow)
}

func (suite *ParserTestSuite) TestParseDelayState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	delay, _ := gmf.modules[0].states["Exact_Delay"].(*DelayState)
	suite.Equal(directTransition("Range_Delay"), delay.transition)

	exact := Exact{Quantity: 1, Unit: "years"}
	suite.Equal(exact, delay.exact)

	delay, _ = gmf.modules[0].states["Range_Delay"].(*DelayState)
	suite.Equal(directTransition("ConditionOnset"), delay.transition)

	rng := Range{Low: 1, High: 3, Unit: "years"}
	suite.Equal(rng, delay.rng)
}

func (suite *ParserTestSuite) TestParseConditionOnsetState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	co, _ := gmf.modules[0].states["ConditionOnset"].(*ConditionOnsetState)
	suite.Equal(directTransition("Encounter"), co.transition)
	suite.Equal("Encounter", co.targetEncounter)
	suite.Equal("condition", co.assignToAttribute)

	codes := []Code{Code{System: "SNOMED-CT", Code: "44054006", Display: "Diabetes mellitus"}}
	suite.Equal(codes, co.codes)
}

func (suite *ParserTestSuite) TestParseEncounterState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	encounter, _ := gmf.modules[0].states["Encounter"].(*EncounterState)
	suite.Equal(directTransition("MedicationOrder"), encounter.transition)
	suite.Equal("ambulatory", encounter.class)
	suite.Equal("condition", encounter.reason)

	codes := []Code{Code{System: "SNOMED-CT", Code: "12345678", Display: "Encounter for problem"}}
	suite.Equal(codes, encounter.codes)
}

func (suite *ParserTestSuite) TestParseMedicationOrderState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	medicationOrder, _ := gmf.modules[0].states["MedicationOrder"].(*MedicationOrderState)
	suite.Equal(directTransition("CarePlanStart"), medicationOrder.transition)
	suite.Equal("Encounter", medicationOrder.targetEncounter)
	suite.Equal("medication", medicationOrder.assignToAttribute)
	suite.Equal("condition", medicationOrder.reason)

	codes := []Code{Code{System: "RxNorm", Code: "123456", Display: "Acetaminophen 325mg [Tylenol]"}}
	suite.Equal(codes, medicationOrder.codes)
}

func (suite *ParserTestSuite) TestParseCarePlanStartState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	careplan, _ := gmf.modules[0].states["CarePlanStart"].(*CarePlanStartState)
	suite.Equal(directTransition("Procedure"), careplan.transition)
	suite.Equal("Encounter", careplan.targetEncounter)
	suite.Equal("careplan", careplan.assignToAttribute)
	suite.Equal("condition", careplan.reason)

	codes := []Code{Code{System: "SNOMED-CT", Code: "987654321", Display: "Examplitis care"}}
	suite.Equal(codes, careplan.codes)

	activities := []Code{
		Code{System: "SNOMED-CT", Code: "987654321", Display: "Examplitis therapy"},
		Code{System: "SNOMED-CT", Code: "987654321", Display: "Examplotomy"},
	}
	suite.Equal(activities, careplan.activities)
}

func (suite *ParserTestSuite) TestParseProcedureState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	procedure, _ := gmf.modules[0].states["Procedure"].(*ProcedureState)
	suite.Equal(directTransition("Condition_End_By_Attribute"), procedure.transition)
	suite.Equal("Encounter", procedure.targetEncounter)
	suite.Equal("condition", procedure.reason)

	codes := []Code{Code{System: "SNOMED-CT", Code: "987654321", Display: "Examplotomy"}}
	suite.Equal(codes, procedure.codes)
}

func (suite *ParserTestSuite) TestParseConditionEndState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	end, _ := gmf.modules[0].states["Condition_End_By_Attribute"].(*ConditionEndState)
	suite.Equal(directTransition("Condition_End_By_Name"), end.transition)
	suite.Equal("condition", end.referencedByAttribute)

	end, _ = gmf.modules[0].states["Condition_End_By_Name"].(*ConditionEndState)
	suite.Equal(directTransition("Condition_End_By_Code"), end.transition)
	suite.Equal("ConditionOnset", end.conditionOnset)

	end, _ = gmf.modules[0].states["Condition_End_By_Code"].(*ConditionEndState)
	suite.Equal(directTransition("Medication_End_By_Attribute"), end.transition)

	codes := []Code{Code{System: "SNOMED-CT", Code: "44054006", Display: "Diabetes mellitus"}}
	suite.Equal(codes, end.codes)
}

func (suite *ParserTestSuite) TestParseMedicationEndState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	end, _ := gmf.modules[0].states["Medication_End_By_Attribute"].(*MedicationEndState)
	suite.Equal(directTransition("Medication_End_By_Name"), end.transition)
	suite.Equal("medication", end.referencedByAttribute)

	end, _ = gmf.modules[0].states["Medication_End_By_Name"].(*MedicationEndState)
	suite.Equal(directTransition("Medication_End_By_Code"), end.transition)
	suite.Equal("MedicationOrder", end.medicationOrder)

	end, _ = gmf.modules[0].states["Medication_End_By_Code"].(*MedicationEndState)
	suite.Equal(directTransition("CarePlan_End_By_Attribute"), end.transition)

	codes := []Code{Code{System: "RxNorm", Code: "123456", Display: "Acetaminophen 325mg [Tylenol]"}}
	suite.Equal(codes, end.codes)
}

func (suite *ParserTestSuite) TestParseCarePlanEndState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	end, _ := gmf.modules[0].states["CarePlan_End_By_Attribute"].(*CarePlanEndState)
	suite.Equal(directTransition("CarePlan_End_By_Name"), end.transition)
	suite.Equal("careplan", end.referencedByAttribute)

	end, _ = gmf.modules[0].states["CarePlan_End_By_Name"].(*CarePlanEndState)
	suite.Equal(directTransition("CarePlan_End_By_Code"), end.transition)
	suite.Equal("CarePlanStart", end.careplan)

	end, _ = gmf.modules[0].states["CarePlan_End_By_Code"].(*CarePlanEndState)
	suite.Equal(directTransition("Exact_Observation"), end.transition)

	codes := []Code{Code{System: "SNOMED-CT", Code: "987654321", Display: "Examplitis care"}}
	suite.Equal(codes, end.codes)
}

func (suite *ParserTestSuite) TestParseObservationState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	observation, _ := gmf.modules[0].states["Exact_Observation"].(*ObservationState)
	suite.Equal(directTransition("Range_Observation"), observation.transition)
	suite.Equal("Encounter", observation.targetEncounter)

	exact := Exact{Quantity: 5, Unit: "mL"}
	suite.Equal(exact, observation.exact)

	codes := []Code{Code{System: "LOINC", Code: "1234-5", Display: "Volume"}}
	suite.Equal(codes, observation.codes)

	observation, _ = gmf.modules[0].states["Range_Observation"].(*ObservationState)
	suite.Equal(directTransition("Exact_Symptom"), observation.transition)
	suite.Equal("Encounter", observation.targetEncounter)

	rng := Range{Low: 2, High: 7, Unit: "mg"}
	suite.Equal(rng, observation.rng)

	codes = []Code{Code{System: "LOINC", Code: "1234-5", Display: "Weight"}}
	suite.Equal(codes, observation.codes)
}

func (suite *ParserTestSuite) TestParseSymptomState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	symptom, _ := gmf.modules[0].states["Exact_Symptom"].(*SymptomState)
	suite.Equal(directTransition("Range_Symptom"), symptom.transition)
	suite.Equal("I'd rather have a bottle in front of me...", symptom.cause)
	exact := Exact{Quantity: 50, Unit: ""}
	suite.Equal(exact, symptom.exact)

	symptom, _ = gmf.modules[0].states["Range_Symptom"].(*SymptomState)
	suite.Equal(directTransition("SetAttributeString"), symptom.transition)
	suite.Equal("...than a frontal lobotomy.", symptom.cause)
	rng := Range{Low: 10, High: 20, Unit: ""}
	suite.Equal(rng, symptom.rng)
}

func (suite *ParserTestSuite) TestParseSetAttributeState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	var sa *SetAttributeState

	suite.NotPanics(func() { sa, _ = gmf.modules[0].states["SetAttributeString"].(*SetAttributeState) })
	suite.Equal(directTransition("SetAttributeNil"), sa.transition)
	suite.Equal("attribute", sa.attribute)
	suite.Equal("string", sa.getStringValue())
	suite.Panics(func() { sa.getFloatValue() })
	suite.Panics(func() { sa.getIntValue() })
	suite.Panics(func() { sa.getBooleanValue() })

	suite.NotPanics(func() { sa, _ = gmf.modules[0].states["SetAttributeNumeric"].(*SetAttributeState) })
	suite.Equal(directTransition("SetAttributeNil"), sa.transition)
	suite.Equal("attribute", sa.attribute)
	suite.Equal(7.1, sa.getFloatValue())
	// If the value is a float, getIntValue() will cast it to an int
	suite.Equal(7, sa.getIntValue())
	suite.Panics(func() { sa.getStringValue() })
	suite.Panics(func() { sa.getBooleanValue() })

	suite.NotPanics(func() { sa, _ = gmf.modules[0].states["SetAttributeBoolean"].(*SetAttributeState) })
	suite.Equal(directTransition("SetAttributeNil"), sa.transition)
	suite.Equal("attribute", sa.attribute)
	suite.Equal(false, sa.getBooleanValue())
	suite.Panics(func() { sa.getStringValue() })
	suite.Panics(func() { sa.getFloatValue() })
	suite.Panics(func() { sa.getIntValue() })
}

func (suite *ParserTestSuite) TestParseSetCounterState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	counter, _ := gmf.modules[0].states["Counter"].(*CounterState)
	suite.Equal(directTransition("Immediate_Death"), counter.transition)
	suite.Equal("attribute", counter.attribute)
	suite.Equal("increment", counter.action)
}
func (suite *ParserTestSuite) TestParseDeathState() {
	gmf := new(GMF)
	gmf.loadModule("../fixtures/states.json")

	death, _ := gmf.modules[0].states["Immediate_Death"].(*DeathState)
	suite.Equal(directTransition("Exact_Death"), death.transition)

	death, _ = gmf.modules[0].states["Exact_Death"].(*DeathState)
	suite.Equal(directTransition("Range_Death"), death.transition)

	exact := Exact{Quantity: 1, Unit: "days"}
	suite.Equal(exact, death.exact)

	death, _ = gmf.modules[0].states["Range_Death"].(*DeathState)
	suite.Equal(directTransition("Terminal"), death.transition)

	rng := Range{Low: 1, High: 2, Unit: "days"}
	suite.Equal(rng, death.rng)
}
func directTransition(nextState string) *DirectTransition {
	return &DirectTransition{
		nextState: nextState,
	}
}

// ============================================================================
// TEST INVALID STATES
// ============================================================================
// This section tests parsing states that we know are invalid and/or incomplete.

// TODO: Validate each individual state type on load. For now, just validating
// aspects common to all states and necessary to process the state.

func (suite *ParserTestSuite) TestParseInvalidStateNoType() {
	gmf := new(GMF)
	err := gmf.loadModule("../fixtures/invalid_states/invalid_state_no_type.json")
	suite.NotNil(err)
	suite.Equal(errors.New("Invalid Module: Invalid State 'Initial': No state type found"), err)
}

func (suite *ParserTestSuite) TestParseInvalidStateNoTransition() {
	gmf := new(GMF)
	err := gmf.loadModule("../fixtures/invalid_states/invalid_state_no_transition.json")
	suite.NotNil(err)
	suite.Equal(errors.New("Invalid Module: Invalid State 'Initial': No valid transition found"), err)
}

func (suite *ParserTestSuite) TestParseInvalidStateUnknownStateType() {
	gmf := new(GMF)
	err := gmf.loadModule("../fixtures/invalid_states/invalid_state_unknown_type.json")
	suite.NotNil(err)
	suite.Equal(errors.New("Invalid Module: Invalid State 'Foo': Unknown state type 'Bar'"), err)
}

func (suite *ParserTestSuite) TestParseInvalidStateInvalidCondition() {
	gmf := new(GMF)
	err := gmf.loadModule("../fixtures/invalid_states/invalid_state_invalid_condition.json")
	suite.NotNil(err)
	suite.Equal(errors.New("Invalid Module: Invalid State 'Guard': No condition type found"), err)
}
