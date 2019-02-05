package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5} 	   // array
	var myslice = []int{1, 2, 3, 4, 5} // slice
	myslice = append(myslice, 6, 7, 1)

	s := myslice[0:3]
	fmt.Println(a)
	fmt.Println(myslice)
	fmt.Println(s)
}
