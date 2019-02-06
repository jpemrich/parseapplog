package main

import (
	"fmt"
	"regexp"

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
	// Testing regular expressions
	r := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})`)
	fmt.Printf("%#v\n", r.FindStringSubmatch(`2015-05-27`))
	fmt.Printf("%#v\n", r.SubexpNames())
	// From docs
	re := regexp.MustCompile("(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)")
	fmt.Println(re.MatchString("Alan Turing"))
	fmt.Printf("%q\n", re.SubexpNames())
	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed))
}
