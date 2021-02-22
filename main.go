package main

import (
	"fmt"
)

var x int = 1 // 00000001
var w int = 2 // 00000010
var r int = 4 // 00000100

func main() {
	bitActive(x | w | r) // 00000111
	bitActive(w | r)     // 00000110
	bitActive(r)
	bitActive(w)
	bitActive(x)

}

func bitActive(res int) {
	fmt.Printf("Num: %d, Bit: %08b\n", res, res)
	var perActive string

	if res&r != 0 {
		perActive += "r"
	}

	if res&w != 0 {
		perActive += "w"
	}

	if res&x != 0 {
		perActive += "x"
	}

	fmt.Printf("Permissions: %s\n", perActive)
}
