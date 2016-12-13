package utils

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ChoiceTestSuite struct {
	suite.Suite
	choices   []Choice
	tolerance float64
}

func TestChoiceTestSuite(t *testing.T) {
	suite.Run(t, new(ChoiceTestSuite))
}

func (c *ChoiceTestSuite) SetupTest() {
	c.tolerance = 0.00000001
	c.choices = []Choice{
		{0.7, "a"},
		{0.2, "b"},
		{0.1, "c"},
	}
}

func (c *ChoiceTestSuite) TestWeightedChoice() {
	choice := WeightedChoice(c.choices)
	c.True(contains(c.choices, choice), "Weighted choice not found in possible choices")
}

func (c *ChoiceTestSuite) TestCleanChoicesNoChoices() {
	c.Panics(func() { cleanChoices([]Choice{}) }, "WeightedChoice: No choices provided")
}

func (c *ChoiceTestSuite) TestCleanChoicesBadWeight() {
	c.choices[1].Weight = -3.0
	c.Panics(func() { cleanChoices(c.choices) }, "WeightedChoice: Choice weight < 0")
}

func (c *ChoiceTestSuite) TestCleanChoiceWeightsLessThanOne() {
	c.choices[2].Weight = 0.05
	choices := cleanChoices(c.choices)
	c.Equal(3, len(choices), "Should still have the same number of choices")
	c.True((0.1-choices[2].Weight) < c.tolerance, "Last choice should be scaled to sum to 1")
}

func (c *ChoiceTestSuite) TestCleanChoiceWeightsGreaterThanOne() {
	c.choices[2].Weight = 0.3
	choices := cleanChoices(c.choices)
	c.Equal(3, len(choices), "Should still have the same number of choices")
	// Due to small floating point errors these values will not be *exactly* the
	// same, so we check accuracy to within a small tolerance.
	c.True((0.1-choices[2].Weight) < c.tolerance, "Last choice should be scaled to sum to 1")
}

func (c *ChoiceTestSuite) TestSortChoices() {
	choices := c.choices
	sort.Sort(SortChoice(c.choices))
	c.True(choices[0].Weight == 0.1, "Choices are not sorted in ascending order [0]")
	c.True(choices[1].Weight == 0.2, "Choices are not sorted in ascending order [1]")
	c.True(choices[2].Weight == 0.7, "Choices are not sorted in ascending order [2]")
}

func (c *ChoiceTestSuite) TestWeightAsString() {
	c.Equal("0.100000", c.choices[2].WeightAsString(), "Conversion to string is incorrect")
}

func contains(choices []Choice, want Choice) bool {
	// Tests if a slice of Choices contains a certain Choice
	for _, el := range choices {
		if el == want {
			return true
		}
	}
	return false
}
