package main

import (
	"fmt"

	"github.com/pborman/getopt/v2"
)

func test(x ...string) string {

	for _, v := range x {
		println(v)
	}
	return "x"
}

func main() {
	// Declare the flags to be used
	helpFlag := getopt.BoolLong("help", 'h', "display help", "More help")
	cmdFlag := getopt.StringLong("command", 'c', "", "the command")

	// Parse the program arguments
	getopt.Parse()
	if *helpFlag {
		getopt.Usage()
	}
	// Get the remaining positional parameters
	test("one", "two", "three")
	if *cmdFlag != "" {
		fmt.Printf("Running command %s", *cmdFlag)
	}
}
