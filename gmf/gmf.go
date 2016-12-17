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

	var module Module
	err = json.Unmarshal(data, &module)
	if err != nil {
		return err
	}

	fmt.Printf("Loaded module '%s'\n", module.Name)
	gmf.modules = append(gmf.modules, module)
	return nil
}

func (gmf *GMF) Process(entity *entity.Entity, time time.Time) {
	// for _, module := range gmf.modules {
	// 	module.Process(entity, time)
	// }
}
