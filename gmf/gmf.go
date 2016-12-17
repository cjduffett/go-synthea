package gmf

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"io/ioutil"

	"github.com/cjduffett/synthea/entity"
)

// GMF is the top-level interface to the Generic Module Framework.
type GMF struct {
	modules []Module
}

// Load loads all the GMF modules found in a given directory.
func (gmf *GMF) Load(moduleDir string) error {
	var err error
	files, err := ioutil.ReadDir(moduleDir)
	if err != nil {
		return err
	}

	// Trim the trailing slash off the directory name (if present)
	// then add it back in, for consistency.
	dir := strings.TrimRight(moduleDir, "/") + "/"

	for _, file := range files {
		if !file.IsDir() {
			err = gmf.loadModule(dir + file.Name())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (gmf *GMF) loadModule(filePath string) error {
	var err error
	if !strings.HasSuffix(filePath, ".json") {
		return errors.New("Not a valid JSON module file")
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// First load the module in its JSON representation
	var jmodule JSONModule
	err = json.Unmarshal(data, &jmodule)
	if err != nil {
		return err
	}

	// Then parse the JSON representation into a concrete Module and States
	module := NewModule(jmodule.Name)

	for name, jsonState := range jmodule.JSONStates {
		state, err := parseState(name, jsonState)
		if err != nil {
			return err
		}
		module.states[name] = state
	}
	gmf.modules = append(gmf.modules, *module)

	fmt.Printf("Loaded module '%s'\n", module.name)
	return nil
}

// Run runs an entity through all of the modules at a given time.
func (gmf *GMF) Run(entity *entity.Entity, time time.Time) {}

func getStateNames(stateMap map[string]JSONState) []string {
	var keys []string
	for key := range stateMap {
		keys = append(keys, key)
	}
	return keys
}
