package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("123")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("String value %d is successfully converted to int", n)
	}

}