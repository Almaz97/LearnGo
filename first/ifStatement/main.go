package main

import (
	"fmt"
	"os"
)

func main() {

	accessForUser := "Almaz"
	usersPassword := "1235"
	username := os.Args[1]
	password := os.Args[2]

	if username == accessForUser && password == usersPassword {
		fmt.Printf("Access provided for %s\n", username)
	} else {
		fmt.Printf("Access denied for %s\n", username)
	}
}
