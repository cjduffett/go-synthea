package utils

import (
	"fmt"
	"math/rand"
	"sort"
)

// DecimalPlaces is the number of decimal places to round all
// floating point numbers to.
var DecimalPlaces = 6

// Choice is an element of a weighted choice array
type Choice struct {
	Weight float64
	Item   interface{}
}

// WeightAsString returns the weight of Choice as a string,
// truncated to 6 decimal places.
func (c *Choice) WeightAsString() string {
	return fmt.Sprintf("%f", c.Weight)
}

// WeightedChoice selected a choice given its probability
// of being selected.
func WeightedChoice(choices []Choice) Choice {
	cleaned := cleanChoices(choices)

	// sort by weights
	sort.Sort(SortChoice(cleaned))

	// pick a random number that will fall in that range
	r := rand.Float64()
	for i := range choices {
		if r < choices[i].Weight {
			return choices[i]
		}
	}
	// If the only choice has weight of 1
	return choices[0]
}

// If weights sum to >1.0, we ignore the remaining choices
// and scale the probability of the last choice in the list.
// If weights are <1.0, we give the last choice an effective
// weight to make the results sum to 1.0.
func cleanChoices(choices []Choice) []Choice {

	if len(choices) == 0 {
		panicOnBadChoice("No choices provided")
	}
	last := len(choices) - 1

	total := 0.0
	for i := range choices {
		weight := choices[i].Weight
		if weight <= 0 {
			panicOnBadChoice("Choice weight <= 0")
		}

		total += weight
		if total > 1 {
			last = i
			break
		}
	}

	// Truncate to a fixed number of decimal places to avoid floating point imprecision
	diff := 1 - total
	choices[last].Weight += diff

	return choices[0 : last+1]
}

func panicOnBadChoice(message string) {
	panic(fmt.Sprintf("WeightedChoice: %s\n", message))
}

// SortChoice implements the sort.Interface for []Choice
// based on the Weight field.
type SortChoice []Choice

func (s SortChoice) Len() int           { return len(s) }
func (s SortChoice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortChoice) Less(i, j int) bool { return s[i].Weight < s[j].Weight }
