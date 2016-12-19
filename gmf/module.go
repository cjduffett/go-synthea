package gmf

import "fmt"

// Module is a GMF module, for example "Diabetes". Each JSON module
// is loaded into this struct for processing by the GMF.
type Module struct {
	name    string
	states  map[string]State
	context *Context
}

// NewModule returns a new initialized GMF module.
func NewModule(name string) *Module {
	return &Module{
		name:    name,
		states:  make(map[string]State),
		context: NewContext(),
	}
}

// Context is the current context of a GMF module, including the current
// state being processed and the module's history of processed states.
type Context struct {
	history      []State
	currentState State
}

// NewContext returns a new initialized module context.
func NewContext() *Context {
	return &Context{
		history:      []State{},
		currentState: nil,
	}
}

// Init initializes (or re-initializes) the module, readying it to
// process a patient's record.
func (m *Module) Init() {
	m.context.history = make([]State, len(m.states))
	initial, ok := m.states["Initial"]
	if !ok {
		panic(fmt.Sprintf("No Initial state found in module %s", m.name))
	}
	m.context.currentState = initial
}

/*

// Process processes the next state(s) in the module until
// a blocking state or the "Terminal" state is reached.
func (m *Module) Process(entity *entity.Entity, time time.Time) {

	// Time is passed as a pointer so Delay states may modify (rewind)
	// the time within this module without modifying the global
	// simulation time.
		for m.context.currentState.process(entity, &time) {
			nextStateName := m.context.currentState.transition()
			nextState, ok := m.states[nextStateName]

			if currentState.name == nextState.name {

			}

			if !ok {
				panic(fmt.Sprintf("Attempted to transition to state '%s': state not found", nextStateName))
			}
			if nextState != m.context.currentState {
				m.context.history = append(m.context.history, m.context.currentState)
			} else {
				// This state loops back to itself. Just upate the timestamps.
			}

			m.context.currentState = nextState
		}
}

// Processed returns true if the module has been fully processed and
// the module has reached a Terminal state.
func (m *Module) Processed() bool {
	return m.currentState.(type) == Terminal
}

*/
