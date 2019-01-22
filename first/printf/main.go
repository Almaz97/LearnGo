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
}
