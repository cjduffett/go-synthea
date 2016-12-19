package gmf

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
	suite.True(isValidUnitOfTime(unit))
	unit = "foo"
	suite.False(isValidUnitOfTime(unit))
}

func (suite *UnitsTestSuite) TestConvertTimeToDuration() {
	var quantity int64
	var unit string

	quantity = 1
	unit = "seconds"
	suite.Equal(time.Second, convertTimeToDuration(quantity, unit))

	quantity = 5
	unit = "minutes"
	suite.Equal(5*time.Minute, convertTimeToDuration(quantity, unit))

	quantity = 3
	unit = "hours"
	suite.Equal(3*time.Hour, convertTimeToDuration(quantity, unit))

	quantity = 4
	unit = "days"
	suite.Equal(time.Hour*24*4, convertTimeToDuration(quantity, unit))

	quantity = 6
	unit = "weeks"
	suite.Equal(time.Hour*24*7*6, convertTimeToDuration(quantity, unit))

	quantity = 10
	unit = "years"
	suite.Equal(time.Hour*24*365*10, convertTimeToDuration(quantity, unit))

	quantity = 1
	unit = "foo"
	suite.Panics(func() { convertTimeToDuration(quantity, unit) })
}
