package sequential

import (
	"fmt"
	"time"

	"github.com/cjduffett/synthea/patient"
)

// Task executes a sequential generation of patients.
type Task struct {
	startDate      time.Time
	endDate        time.Time
	timeStep       int
	numToGenerate  int
	livingPopCount int
	deadPopCount   int
}

// NewTask returns a new sequential run to execute.
func NewTask(numToGenerate int) *Task {
	now := time.Now()
	return &Task{
		endDate:        now,
		startDate:      now.AddDate(-100, 0, 0),
		timeStep:       7,
		numToGenerate:  numToGenerate,
		livingPopCount: 0,
		deadPopCount:   0,
	}

	// TODO: Track patient statistics
}

// Run executes a sequential Synthea generation.
func (task *Task) Run() {
	if task.numToGenerate == 0 {
		// The world is not initialized yet
		panic("World not initialized")
	}

	fmt.Printf("Generating %d patients...\n", task.numToGenerate)
	task.runRandom()
	// TODO: support multithreading

	// load modules

	// execute modules

	// export
}

func (task *Task) runRandom() {

	// TODO: randomize seed?

	for task.livingPopCount < task.numToGenerate {
		// create a new patient
		fmt.Printf("Patient... %d\n", task.livingPopCount)
		patient.NewPatient(task.startDate, task.endDate)
		task.livingPopCount++
	}
}
