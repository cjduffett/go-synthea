package main

import (
	"fmt"
	"os"

	"github.com/cjduffett/synthea/cli"
)

func main() {

	// usage guide
	if len(os.Args) == 1 {
		fmt.Println("usage: synthea <command> [<args>]")
		fmt.Println("The most commonly used commands are: ")
		fmt.Println(" sequential   Sequentially generate patients ")
		fmt.Println(" graphviz     Create a graphical vizualization of synthea modules ")
		fmt.Println(" story        Create a \"story\" of a patient's life ")
		fmt.Println(" new          Create a new generic module ")
		return
	}

	// sub commands
	cli.ParseSubCommands(os.Args[1], os.Args[2:])

	// read in config

	// build world

	// init simulation

	// switch simulation type

	// thread it

	// export (fhir, CCDA, html)

	// generate summary

	// graphviz?

}
