package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Task feet convert to meters
	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)
	const feetToMeters = 0.3048
	const feetToYard = 0.3333
	meters := feet * feetToMeters
	yards := feet * feetToYard
	fmt.Printf("%g feet %g meters: \n", feet, meters)
	fmt.Printf("%g feet %g meters: \n", feet, yards)

	// Task celsius convert to fahrenheit

	//arg := os.Args[1]
	//cel, _ := strconv.ParseFloat(arg, 64)
	//fah := cel*1.8 + 32
	//fmt.Printf("%g celsius - %g fahrenheit\n", cel, fah)
}
