package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	name = os.Args[1]
	fmt.Println("Hello ", name)
	name = os.Args[2]
	fmt.Println("Hello ", name)
	name = os.Args[3]
	fmt.Println("Hello ", name)
	name = os.Args[4]
	fmt.Println("Hello ", name)

}
