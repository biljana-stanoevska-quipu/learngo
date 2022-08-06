package main

import (
	"fmt"
)

func main() {

	type person struct {
		name string
		height float64
	}

	// var p person
	// p := new(person)

	p := person{
		name: "biljana",
		height: 168.0,
	}

	fmt.Println(p)
	fmt.Println("name is", p.name)
	fmt.Println("height is", p.height)
	p.height = 168.5
	fmt.Println("height is", p.height)
}