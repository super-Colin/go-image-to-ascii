package main

import (
	"flag"
	"fmt"
)

type scoFlag interface {
	init()
	validate()
}

type scoFlag_int struct {
	value        int         // The value of the flag
	flagName     string      // The name of the flag on the CLI
	description  string      // What does this flag do?
	min          int         // -1 if not used, inclusive
	max          int         // -1 if not used, inclusive
	defaultVal   int         // Fallback value
	validateFunc func() bool // Validation
	initFunc     func()      // Set .value to pointer of CLI flag value
}

func MakeIntFlag(cliName, description string, defaultVal, min, max int) scoFlag_int {
	var cliDescription = fmt.Sprintf("%d~%d?=%d; ", min, max, defaultVal)
	f := scoFlag_int{
		value:       0,
		flagName:    cliName,
		description: cliDescription + description,
		min:         min,
		max:         max,
		defaultVal:  defaultVal,
	}
	f.initFunc = f.init_default
	f.validateFunc = f.validate_default

	return f
}
func (sf *scoFlag_int) init() {
	sf.initFunc()
}
func (sf *scoFlag_int) validate() {
	sf.validateFunc()
}
func (sf *scoFlag_int) init_default() {
	flag.IntVar(&sf.value, sf.flagName, sf.defaultVal, sf.description)
}
func (sf *scoFlag_int) validate_default() bool {
	if sf.min != -1 && sf.value < sf.min || sf.max != -1 && sf.value > sf.max {
		fmt.Printf("%s value is out of bounds: %d~%d", sf.flagName, sf.min, sf.max)
		return false
	}
	return true
}
