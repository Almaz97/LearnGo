package main

import (
	"fmt"
	"reflect"
)

func main() {
	speed := 100
	force := 2.5

	speed = force * speed
	fmt.Println(reflect.TypeOf(force))
	fmt.Println(speed)
}
