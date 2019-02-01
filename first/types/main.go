package main

import "fmt"

func main() {
	type Duration int64
	var ns Duration = 10000
	var ms float64 = 10

	fmt.Printf("ms: %f, ns: %d\n", ms, ns)
	res := ns + Duration(ms)
	fmt.Printf("res: %d\n", res)
}
