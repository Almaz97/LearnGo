package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var msg = os.Args[1]
	var l = len(msg)
	fmt.Println(strings.Repeat("!", l) + strings.ToUpper(msg) + strings.Repeat("!", l))

}
