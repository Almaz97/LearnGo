package main

import (
	"fmt"
	"math"
)

func main() {

	var n int8 = math.MaxInt8
	fmt.Println("max int 8: ", n)
	fmt.Println("max int 8 + 1: ", n+1)

	n = math.MinInt8
	fmt.Println("max int 8: ", n)
	fmt.Println("max int 8 - 1: ", n-1)

	var un uint8
	fmt.Println("max uint8: ", un)
	fmt.Println("max uint8 - 1: ", un-1)

	f := float32(math.MaxFloat32)
	fmt.Println("max float :", f*1.2)
}
