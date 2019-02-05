package main

import "fmt"

func main() {
	const (
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

	const (
		Est = -(5 + iota)
		_
		MST
		PST
	)
	fmt.Println(Est, MST, PST)

}