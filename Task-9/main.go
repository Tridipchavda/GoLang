package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define subcommands
	subcommand1 := flag.NewFlagSet("subcommand1", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("subcommand2", flag.ExitOnError)

	// Define flags for subcommand1
	flagName := subcommand1.String("name", "", "Specify name for subcommand1")

	// Define flags for subcommand2
	flagAge := subcommand2.Int("age", 0, "Specify age for subcommand2")

	// Parse the command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Please specify a subcommand.")
		os.Exit(1)
	}

	//Check which subcommand is enter by user
	switch os.Args[1] {
	case "subcommand1":
		// Parse subcommand to get the flag Name
		subcommand1.Parse(os.Args[2:])
		fmt.Println("Executing subcommand1")
		fmt.Println("Name:", *flagName)
	case "subcommand2":
		// Parse subcommand to get the flag Age
		subcommand2.Parse(os.Args[2:])
		fmt.Println("Executing subcommand2")
		fmt.Println("Age:", *flagAge)
	default:
		// Default message If some different command is entered by user
		fmt.Println("Invalid subcommand.")
		os.Exit(1)
	}

}
