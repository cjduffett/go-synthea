package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/cjduffett/synthea/world"
)

// ParseSubCommands parses the specified command and its arguments
func ParseSubCommands(cmd string, args []string) {

	switch cmd {

	case "sequential":
		// sequential [options]
		// -n         Number of patients to generate (default 100)

		// TODO: Additional config
		// -config    Path to custom synthea.yml (default is at config/synthea.yml)
		// -demo      Provide demographic data (see towns.json)
		// -thread    Multithread the sequential generation

		sequentialCommand := flag.NewFlagSet("sequential", flag.ExitOnError)
		numPatients := sequentialCommand.Int("n", 100, "The number of patients to generate ")

		// parse the args
		sequentialCommand.Parse(args)
		if sequentialCommand.Parsed() {
			world.NewSequentialTask(*numPatients).Run()
		}

	default:
		notImplemented(cmd)
	}
}

// TODO: Additional sub commands
// graphviz [module_name]
//graphvizCommand := flag.NewFlagSet("graphviz", flag.ExitOnError)
//
// story [patient_id]
//storyCommand := flag.NewFlagSet("story", flag.ExitOnError)
//
// new [module_name]
//newModuleCommand := flag.NewFlagSet("new", flag.ExitOnError)

func invalidArgs(cmd string, err error) {
	fmt.Printf("Invalid arguments for command %s.\n", cmd)
	fmt.Println(err)
	os.Exit(2)
}

func notImplemented(cmd string) {
	fmt.Printf("%s is not a valid command.\n", cmd)
	os.Exit(2)
}
