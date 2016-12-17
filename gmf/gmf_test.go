package gmf

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type GMFTestSuite struct {
	suite.Suite
}

func TestGMFTestSuite(t *testing.T) {
	suite.Run(t, new(GMFTestSuite))
}

func (suite *GMFTestSuite) TestLoadEmptyModule() {
	gmf := new(GMF)
	var err error
	suite.NotPanics(func() { err = gmf.loadModule("../fixtures/gmf/empty_module.json") })
	suite.Nil(err)
	suite.Equal(1, len(gmf.modules))

	module := gmf.modules[0]
	suite.Equal("Empty Module", module.name)
	keys := getKeys(module.states)
	suite.Equal(0, len(keys))
}

func (suite *GMFTestSuite) TestLoadBasicModule() {
	gmf := new(GMF)
	var err error
	suite.NotPanics(func() { err = gmf.loadModule("../fixtures/gmf/basic_module.json") })
	suite.Nil(err)
	suite.Equal(1, len(gmf.modules))

	module := gmf.modules[0]
	suite.Equal("Basic Module", module.name)
	keys := getKeys(module.states)
	suite.Equal(2, len(keys))
}

func (suite *GMFTestSuite) TestLoadModules() {
	gmf := new(GMF)
	var err error
	suite.NotPanics(func() { err = gmf.Load("../fixtures/gmf") })
	suite.Nil(err)
	suite.Equal(2, len(gmf.modules))

	module := gmf.modules[0]
	suite.Equal("Basic Module", module.name)
	keys := getKeys(module.states)
	suite.Equal(2, len(keys))

	module = gmf.modules[1]
	suite.Equal("Empty Module", module.name)
	keys = getKeys(module.states)
	suite.Equal(0, len(keys))
}

func getKeys(stateMap map[string]State) []string {
	var keys []string
	for key := range stateMap {
		keys = append(keys, key)
	}
	return keys
}
