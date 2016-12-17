package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UnitsTestSuite struct {
	suite.Suite
}

func TestUnitsTestSuite(t *testing.T) {
	suite.Run(t, new(UnitsTestSuite))
}

func (suite *UnitsTestSuite) TestIsValidUnitOfTime() {
	unit := "hours"
	suite.True(IsValidUnitOfTime(unit))
	unit = "foo"
	suite.False(IsValidUnitOfTime(unit))
}

func (suite *UnitsTestSuite) TestIsValidUnitOfAge() {
	unit := "years"
	suite.True(IsValidUnitOfAge(unit))
	unit = "bar"
	suite.False(IsValidUnitOfAge(unit))
}
