package main

import (
	"fmt"
)

func main() {

	// friends := make([]string, 5, 10)  // type, length, capacity (max size of the slice)

	friends := []string{
		"ivana",
		"angela",
		"ninka"}

	fmt.Printf("Length of slice is %d and capacity is %d.\n", len(friends), cap(friends))

	for _, i := range friends {
		fmt.Println(i)
	}

	// initial capacity is 12 = 2 * 6
	myIntSlice := []int{1, 2, 3, 100, 1000, 1000000}
	fmt.Println(myIntSlice)

	myIntSlice[0] = 10
	fmt.Println(myIntSlice)

	sliceOfSlice := myIntSlice[1:4] // 1 is included, 4 is not included
	fmt.Println(sliceOfSlice)

	for i := 1; i < 10; i++ {
		myIntSlice = append(myIntSlice, i)
		fmt.Println(myIntSlice)
		// doubles the capacity of the background array when we are out of indexes
		fmt.Printf("Length of slice is %d and capacity is %d.\n", len(myIntSlice), cap(myIntSlice))
	}


	fmt.Println("myIntSlice", myIntSlice)
	newSlice := []int{11, 22, 33}
	fmt.Println("newSlice", newSlice)
	myIntSlice = append(myIntSlice, newSlice...)
	fmt.Println("myIntSlice", myIntSlice)




}