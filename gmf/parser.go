package gmf

import (
	"errors"
	"fmt"
)

// JSONModule is the JSON representation of a module.
type JSONModule struct {
	Name       string               `json:"name"`
	JSONStates map[string]JSONState `json:"states"`
}

// JSONState is a newly unmarshalled JSON state that must be
// parsed into a concrete State type. We let the unmarshaller do
// most of the work and only chose a subset of the unmarshalled
// properties (that aren't nil) to build a State based on the state's
// Type. Since we cannot know the type of state we are about to
// unmarshal until runtime this is the simplest approach.
type JSONState struct {
	Type                  string            `json:"type"`
	TargetEncounter       string            `json:"target_encounter"`
	EncounterClass        string            `json:"encounter_class"`
	Reason                string            `json:"reason"`
	ConditionOnset        string            `json:"condition_onset"`
	MedicationOrder       string            `json:"medication_order"`
	CarePlan              string            `json:"careplan"`
	Attribute             string            `json:"attribute"`
	Value                 interface{}       `json:"value"`
	Action                string            `json:"action"`
	AssignToAttribute     string            `json:"assign_to_attribute"`
	ReferencedByAttribute string            `json:"referenced_by_attribute"`
	Exact                 Exact             `json:"exact"`
	Range                 Range             `json:"range"`
	Unit                  string            `json:"unit"`
	Codes                 []Code            `json:"codes"`
	Activities            []Code            `json:"activities"`
	Allow                 JSONCondition     `json:"allow"`
	Wellness              bool              `json:"wellness"`
	Symptom               string            `json:"symptom"`
	Cause                 string            `json:"cause"`
	DirectTransition      string            `json:"direct_transition"`
	DistributedTransition []Distribution    `json:"distributed_transition"`
	ConditionalTransition []JSONConditional `json:"conditional_transition"`
	ComplexTransition     []JSONComplex     `json:"complex_transition"`
}

// JSONCondition is a newly unmarshalled logical condition that must be
// parsed into a concrete Condition type. We let the unmarshaller do most
// of the work and only chose the subset of the unmarshalled properties
// (that aren't nil) to build a Condition object based on the condition's
// ConditionType. Same approach as JSONState.
type JSONCondition struct {
	ConditionType         string          `json:"condition_type"`
	Conditions            []JSONCondition `json:"conditions"`
	Condition             *JSONCondition  `json:"condition"`
	Operator              string          `json:"operator"`
	Quantity              float64         `json:"quantity"`
	Value                 interface{}     `json:"value"`
	Unit                  string          `json:"unit"`
	Gender                string          `json:"gender"`
	Year                  int             `json:"year"`
	Category              string          `json:"category"`
	Race                  string          `json:"race"`
	Symptom               string          `json:"symptom"`
	Codes                 []Code          `json:"codes"`
	ReferencedByAttribute string          `json:"referenced_by_attribute"`
	Name                  string          `json:"name"`
	Attribute             string          `json:"attribute"`
	Minimum               int             `json:"minimum"`
	Maximum               int             `json:"maximum"`
}

// JSONConditional is a newly unmarshalled Conditional element of a
// conditional transition. The Conditional type can't be used directly
// because of the added complexity of unmarshalling the Condition.
type JSONConditional struct {
	Condition  JSONCondition `json:"condition"`
	Transition string        `json:"transition"`
}

// JSONComplex is a newly unmarshalled Complex element of a
// conditional transition. The Complex type can't be used directly
// because of the added complexity of unmarshalling the Condition.
type JSONComplex struct {
	Condition     JSONCondition  `json:"condition"`
	Distributions []Distribution `json:"distributions"`
}

// parseState parses a JSONState into a GMF State based on the state's type.
func parseState(stateName string, jsonState JSONState) (state State, err error) {
	// There are many ways parsing a state can fail. All component
	// parsing methods simply panic when they encounter an error.
	// All parsing errors are caught here.
	defer func() {
		if r := recover(); r != nil {
			// find out exactly what the error was and set err
			switch t := r.(type) {
			case string:
				err = fmt.Errorf("Invalid State '%s': %s", stateName, t)
			case error:
				err = fmt.Errorf("Invalid State '%s': %s", stateName, t.Error())
			default:
				err = errors.New("Unknown panic")
			}
			state = nil
		}
	}()

	if jsonState.Type == "" {
		panic("No state type found")
	}

	var transition Transition
	if jsonState.Type != "Terminal" {
		transition = parseTransition(jsonState)
	}

	switch jsonState.Type {
	case "Terminal":
		state = parseTerminalState()
	case "Initial":
		state = parseInitialState(transition)
	case "Simple":
		state = parseSimpleState(transition)
	case "Guard":
		state = parseGuardState(jsonState, transition)
	case "Delay":
		state = parseDelayState(jsonState, transition)
	case "Encounter":
		state = parseEncounterState(jsonState, transition)
	case "Procedure":
		state = parseProcedureState(jsonState, transition)
	case "ConditionOnset":
		state = parseConditionOnsetState(jsonState, transition)
	case "ConditionEnd":
		state = parseConditionEndState(jsonState, transition)
	case "MedicationOrder":
		state = parseMedicationOrderState(jsonState, transition)
	case "MedicationEnd":
		state = parseMedicationEndState(jsonState, transition)
	case "CarePlanStart":
		state = parseCarePlanStartState(jsonState, transition)
	case "CarePlanEnd":
		state = parseCarePlanEndState(jsonState, transition)
	case "Symptom":
		state = parseSymptomState(jsonState, transition)
	case "Observation":
		state = parseObservationState(jsonState, transition)
	case "SetAttribute":
		state = parseSetAttributeState(jsonState, transition)
	case "Counter":
		state = parseCounterState(jsonState, transition)
	case "Death":
		state = parseDeathState(jsonState, transition)
	default:
		panic(fmt.Sprintf("Unknown state type '%s'", jsonState.Type))
	}
	return state, nil
}

// parseTransition parses the transition out of a JSONState and returns
// a Transition. The first transition type found is the one used.
func parseTransition(state JSONState) Transition {

	if state.DirectTransition != "" {
		return &DirectTransition{
			nextState: state.DirectTransition,
		}
	}

	if len(state.DistributedTransition) > 0 {
		return &DistributedTransition{
			distributions: state.DistributedTransition,
		}
	}

	if len(state.ConditionalTransition) > 0 {
		return &ConditionalTransition{
			conditionals: parseConditionalTransition(state.ConditionalTransition),
		}
	}

	if len(state.ComplexTransition) > 0 {
		return &ComplexTransition{
			transitions: parseComplexTransition(state.ComplexTransition),
		}
	}
	panic("No valid transition found")
}

// parseConditionalTransition handles the added complexity of parsing a conditional
// with/without a Condition.
func parseConditionalTransition(jsonConditionals []JSONConditional) []Conditional {
	conditionals := make([]Conditional, len(jsonConditionals))
	for i, jcond := range jsonConditionals {
		var condition Condition
		if jcond.Condition.ConditionType == "" && i == len(jsonConditionals)-1 {
			// The last conditional may omit a condition and specify just a transition
			condition = nil
		} else {
			condition = parseCondition(jcond.Condition)
		}
		conditionals[i] = Conditional{
			Condition: condition,
			NextState: jcond.Transition,
		}
	}
	return conditionals
}

// parseComplexTransition handles the added complexity of parsing a complex
// with/without a Condition.
func parseComplexTransition(jsonComplexes []JSONComplex) []Complex {
	complexes := make([]Complex, len(jsonComplexes))
	for i, jcomplex := range jsonComplexes {
		var condition Condition
		if jcomplex.Condition.ConditionType == "" && i == len(jsonComplexes)-1 {
			// The last conditional may omit a condition and specify just a transition
			condition = nil
		} else {
			condition = parseCondition(jcomplex.Condition)
		}
		complexes[i] = Complex{
			Condition:     condition,
			Distributions: jcomplex.Distributions,
		}
	}
	return complexes
}

// parseCondition parses a JSONCondition into a Condition based
// on JSONCondition.ConditionType.
func parseCondition(jsonCondition JSONCondition) Condition {

	switch jsonCondition.ConditionType {
	case "Gender":
		return &GenderCondition{
			gender: jsonCondition.Gender,
		}
	case "Age":
		return &AgeCondition{
			operator: jsonCondition.Operator,
			quantity: jsonCondition.Quantity,
			unit:     jsonCondition.Unit,
		}
	case "Date":
		return &DateCondition{
			operator: jsonCondition.Operator,
			year:     jsonCondition.Year,
		}
	case "Socioeconomic Status":
		return &SocioStatusCondition{
			category: jsonCondition.Category,
		}
	case "Race":
		return &RaceCondition{
			race: jsonCondition.Race,
		}
	case "Symptom":
		// The unmarshaller will already have panicked if the value was not numeric.
		symptomValue, _ := jsonCondition.Value.(float64)
		return &SymptomCondition{
			symptom:  jsonCondition.Symptom,
			operator: jsonCondition.Operator,
			value:    symptomValue,
		}
	case "Observation":
		// The unmarshaller will already have panicked if the value was not numeric.
		observationValue, _ := jsonCondition.Value.(float64)
		return &ObservationCondition{
			referencedByAttribute: jsonCondition.ReferencedByAttribute,
			codes:    jsonCondition.Codes,
			operator: jsonCondition.Operator,
			value:    observationValue,
		}
	case "Active Condition":
		return &ActiveCondition{
			referencedByAttribute: jsonCondition.ReferencedByAttribute,
			codes: jsonCondition.Codes,
		}
	case "Active Medication":
		return &ActiveMedication{
			referencedByAttribute: jsonCondition.ReferencedByAttribute,
			codes: jsonCondition.Codes,
		}
	case "Active CarePlan":
		return &ActiveCarePlan{
			referencedByAttribute: jsonCondition.ReferencedByAttribute,
			codes: jsonCondition.Codes,
		}
	case "PriorState":
		return &PriorStateCondition{
			name: jsonCondition.Name,
		}
	case "Attribute":
		return &AttributeCondition{
			attribute: jsonCondition.Attribute,
			operator:  jsonCondition.Operator,
			value:     jsonCondition.Value,
		}
	case "And":
		return &AndCondition{
			conditions: parseGroupedConditions(jsonCondition.Conditions),
		}
	case "Or":
		return &OrCondition{
			conditions: parseGroupedConditions(jsonCondition.Conditions),
		}
	case "At Least":
		return &AtLeastCondition{
			minimum:    jsonCondition.Minimum,
			conditions: parseGroupedConditions(jsonCondition.Conditions),
		}
	case "At Most":
		return &AtMostCondition{
			maximum:    jsonCondition.Maximum,
			conditions: parseGroupedConditions(jsonCondition.Conditions),
		}
	case "Not":
		return &NotCondition{
			condition: parseCondition(jsonCondition),
		}
	case "True":
		return &TrueCondition{}
	case "False":
		return &FalseCondition{}
	default:
		if jsonCondition.ConditionType == "" {
			panic("No condition type found")
		}
		panic(fmt.Sprintf("Unknown condition type '%s'", jsonCondition.ConditionType))
	}
}

// parseGroupedConditions parses a slice of jsonConditions into a
// slice of GMF Conditions that can be evaulated logically.
func parseGroupedConditions(jsonConditions []JSONCondition) []Condition {
	conditions := make([]Condition, len(jsonConditions))
	for i, jcond := range jsonConditions {
		conditions[i] = parseCondition(jcond)
	}
	return conditions
}

func parseTerminalState() *TerminalState {
	return &TerminalState{}
}

func parseInitialState(transition Transition) *InitialState {
	return &InitialState{
		transition: transition,
	}
}

func parseSimpleState(transition Transition) *SimpleState {
	return &SimpleState{
		transition: transition,
	}
}

func parseGuardState(jsonState JSONState, transition Transition) *GuardState {
	return &GuardState{
		allow:      parseCondition(jsonState.Allow),
		transition: transition,
	}
}

func parseDelayState(jsonState JSONState, transition Transition) *DelayState {
	return &DelayState{
		exact:      jsonState.Exact,
		rng:        jsonState.Range,
		transition: transition,
	}
}

func parseEncounterState(jsonState JSONState, transition Transition) *EncounterState {
	return &EncounterState{
		wellness:   jsonState.Wellness,
		class:      jsonState.EncounterClass,
		reason:     jsonState.Reason,
		codes:      jsonState.Codes,
		transition: transition,
	}
}

func parseProcedureState(jsonState JSONState, transition Transition) *ProcedureState {
	return &ProcedureState{
		targetEncounter: jsonState.TargetEncounter,
		reason:          jsonState.Reason,
		codes:           jsonState.Codes,
		transition:      transition,
	}
}
func parseConditionOnsetState(jsonState JSONState, transition Transition) *ConditionOnsetState {
	return &ConditionOnsetState{
		targetEncounter:   jsonState.TargetEncounter,
		assignToAttribute: jsonState.AssignToAttribute,
		reason:            jsonState.Reason,
		codes:             jsonState.Codes,
		transition:        transition,
	}
}

func parseConditionEndState(jsonState JSONState, transition Transition) *ConditionEndState {
	return &ConditionEndState{
		referencedByAttribute: jsonState.ReferencedByAttribute,
		conditionOnset:        jsonState.ConditionOnset,
		codes:                 jsonState.Codes,
		transition:            transition,
	}
}

func parseMedicationOrderState(jsonState JSONState, transition Transition) *MedicationOrderState {
	return &MedicationOrderState{
		targetEncounter:   jsonState.TargetEncounter,
		assignToAttribute: jsonState.AssignToAttribute,
		reason:            jsonState.Reason,
		codes:             jsonState.Codes,
		transition:        transition,
	}
}

func parseMedicationEndState(jsonState JSONState, transition Transition) *MedicationEndState {
	return &MedicationEndState{
		referencedByAttribute: jsonState.ReferencedByAttribute,
		medicationOrder:       jsonState.MedicationOrder,
		codes:                 jsonState.Codes,
		transition:            transition,
	}
}

func parseCarePlanStartState(jsonState JSONState, transition Transition) *CarePlanStartState {
	return &CarePlanStartState{
		targetEncounter:   jsonState.TargetEncounter,
		assignToAttribute: jsonState.AssignToAttribute,
		reason:            jsonState.Reason,
		codes:             jsonState.Codes,
		activities:        jsonState.Activities,
		transition:        transition,
	}
}

func parseCarePlanEndState(jsonState JSONState, transition Transition) *CarePlanEndState {
	return &CarePlanEndState{
		referencedByAttribute: jsonState.ReferencedByAttribute,
		careplan:              jsonState.CarePlan,
		codes:                 jsonState.Codes,
		transition:            transition,
	}
}

func parseSymptomState(jsonState JSONState, transition Transition) *SymptomState {
	return &SymptomState{
		symptom:    jsonState.Symptom,
		cause:      jsonState.Cause,
		exact:      jsonState.Exact,
		rng:        jsonState.Range,
		transition: transition,
	}
}

func parseObservationState(jsonState JSONState, transition Transition) *ObservationState {
	return &ObservationState{
		targetEncounter: jsonState.TargetEncounter,
		exact:           jsonState.Exact,
		rng:             jsonState.Range,
		unit:            jsonState.Unit,
		codes:           jsonState.Codes,
		transition:      transition,
	}
}

func parseSetAttributeState(jsonState JSONState, transition Transition) *SetAttributeState {
	return &SetAttributeState{
		attribute:  jsonState.Attribute,
		value:      jsonState.Value,
		transition: transition,
	}
}

func parseCounterState(jsonState JSONState, transition Transition) *CounterState {
	return &CounterState{
		attribute:  jsonState.Attribute,
		action:     jsonState.Action,
		transition: transition,
	}
}

func parseDeathState(jsonState JSONState, transition Transition) *DeathState {
	return &DeathState{
		exact:      jsonState.Exact,
		rng:        jsonState.Range,
		transition: transition,
	}
}
