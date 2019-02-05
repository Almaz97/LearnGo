package main

import "fmt"

type gram float64
type ounce float64

func main() {

	type (
		Gram int
		Kilogram Gram
		Ton Kilogram
	)

	var (
		salt Gram = 100
		apples Kilogram = 5
		truck Ton = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)

	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Gram(int(apples)))

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)
	//type some = byte
}
