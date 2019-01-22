package main

import "fmt"

func main() {
	var brand = "Google"
	fmt.Printf("%q\n", brand)
	fmt.Printf("%s\n", brand)

	var ops = 2350
	var ok = 543
	var fail = 433

	fmt.Printf(
		"total: %d, \nsucces: %d / %d\n",
		ops, ok, fail,
	)

	// Escape sequences
	fmt.Println("hi\\n\"hi\"")
	fmt.Println("hi\"This message should appear in double quotes\"")

	// Printing types
	var off bool
	fmt.Printf("%T\n", brand)
	fmt.Printf("%T\n", ops)
	fmt.Printf("%T\n", off)

	// %v in printf
	var (
		planet   = "Venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)
	fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG!\n", planet, distance)
}
