package main

import "fmt"

func main() {
	city := "Limbo"

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	}

	i := 1000
	switch {
	case i > 0:
		fmt.Printf("%d is greater than 0\n", i)
	case i < 0:
		fmt.Printf("%d is less than 0\n", i)
	default:
		fmt.Printf("%d is equal to 0\n", i)
	}

	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number")
	}

}