package utils

import (
	"testing"
	"time"

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

func (suite *UnitsTestSuite) TestConvertTimeToDuration() {
	var unit string
	unit = "seconds"
	suite.Equal(time.Second, ConvertTimeToDuration(unit))
	unit = "minutes"
	suite.Equal(time.Minute, ConvertTimeToDuration(unit))
	unit = "hours"
	suite.Equal(time.Hour, ConvertTimeToDuration(unit))
	unit = "days"
	suite.Equal(time.Hour*24, ConvertTimeToDuration(unit))
	unit = "weeks"
	suite.Equal(time.Hour*24*7, ConvertTimeToDuration(unit))
	unit = "years"
	suite.Equal(time.Hour*24*365, ConvertTimeToDuration(unit))
	unit = "foo"
	suite.Panics(func() { ConvertTimeToDuration(unit) })
}
