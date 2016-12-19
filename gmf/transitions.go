package gmf

import (
	"time"

	"github.com/cjduffett/synthea/entity"
)

// Transition is an interface for all transition types.
type Transition interface {
	follow(entity *entity.Entity, time time.Time) string
}

// DirectTransition transitions directly to the next named state.
type DirectTransition struct {
	nextState string
}

func (dt *DirectTransition) follow(entity *entity.Entity, time time.Time) string {
	return dt.nextState
}

// Conditional maps a logical condition to a state name.
type Conditional struct {
	Condition Condition `json:"condition"`
	NextState string    `json:"transition"`
}

// ConditionalTransition transitions to a state depending on the
// logical conditions provided. If no condition is satisfied,
// ConditionalTransition transitions to the "Terminal" state.
type ConditionalTransition struct {
	conditionals []Conditional
}

// NewConditionalTransition creates a new conditional transition.
func NewConditionalTransition(conditionals []Conditional) Transition {
	return &ConditionalTransition{
		conditionals: conditionals,
	}
}

func (ct *ConditionalTransition) follow(entity *entity.Entity, time time.Time) string {
	for _, conditional := range ct.conditionals {
		if conditional.Condition.test(entity, time) {
			return conditional.NextState
		}
	}
	return "Terminal"
}

// Distribution maps a weight (0 < weight <= 1) to a state name.
type Distribution struct {
	Distribution float64 `json:"distribution"`
	Transition   string  `json:"transition"`
}

// DistributedTransition transitions to each state proportional to
// the distributions provided.
type DistributedTransition struct {
	distributions []Distribution
}

// NewDistributedTransition creates a new distributed transition.
func NewDistributedTransition(distributions []Distribution) Transition {
	return &DistributedTransition{
		distributions: distributions,
	}
}

func (t *DistributedTransition) follow(entity *entity.Entity, time time.Time) string {
	// TODO: Pick a distributed transition
	return ""
}

// Complex maps a logical condition to a series of distributions.
type Complex struct {
	Condition     Condition      `json:"condition"`
	Distributions []Distribution `json:"distributions"`
}

// ComplexTransition transitions to a state depending on both the logical conditions
// provided and the distributions that match those logical conditions.
type ComplexTransition struct {
	transitions []Complex
}

// NewComplexTransition creates a new complex transition.
func NewComplexTransition(transitions []Complex) Transition {
	return &ComplexTransition{
		transitions: transitions,
	}
}

func (t *ComplexTransition) follow(entity *entity.Entity, time time.Time) string {
	// TODO: Pick a complex transition
	return ""
}
