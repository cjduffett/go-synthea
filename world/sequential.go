package world

import (
	"fmt"
	"time"
)

// SequentialTask executes a sequential generation of patients.
type SequentialTask struct {
	startDate      time.Time
	endDate        time.Time
	timeStep       int
	numToGenerate  int
	livingPopCount int
	deadPopCount   int
}

// NewSequentialTask returns a new sequential run to execute.
func NewSequentialTask(numToGenerate int) *SequentialTask {
	now := time.Now()
	return &SequentialTask{
		endDate:       now,
		startDate:     now.AddDate(-100, 0, 0),
		timeStep:      7,
		numToGenerate: numToGenerate,
	}

	// TODO: Track patient statistics
}

// Run executes a sequential Synthea generation.
func (s *SequentialTask) Run() {
	if s.numToGenerate == 0 {
		// The world is not initialized yet
		panic("World not initialized")
	}

	fmt.Printf("Generating %d patients...\n", s.numToGenerate)

	// TODO: support multithreading

	// load modules

	// execute modules

	// export
}
