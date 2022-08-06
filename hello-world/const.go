package main

import (
	"fmt"
)

func main() {
	const c = 300000
	fmt.Println("The light travels with", c, "km per second")

	// c = 200  // cannot assign to c (untyped int constant 300000)
}