package main

import "fmt"

func main() {

	// Rules: can't initialize const equating a var, or some other mutable value
	const (
		min = 250
		max
	)
	fmt.Println(max, min)
	fmt.Printf("%T, %T\n", max, min)

	const (
		some1 = 1.2
		some2 = some1 * 8
	)

	fmt.Println(some1, some2)


	//var a = 35
	//var b = a * 2.3

	var i int = min
	var f float64 = min
	var b byte = min
	var j int32 = min
	var r rune = min
	fmt.Println(i, f, b, j, r)

}