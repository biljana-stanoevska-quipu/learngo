/*
Boolean operator
>
<
==
>=
<=
!
&&
||
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {

	if studyMins, sleepMins := 100, 50; studyMins > sleepMins {
		fmt.Println("You study more than you sleep")
	} else if studyMins == sleepMins {
		fmt.Println("You sleep enough")
	} else {
		fmt.Println("You sleep mora than you study.")
	}

	// studyMins, sleepMins gets GC after the block finishes

	switch "biljana" {
	case "Biljana":
		fmt.Println("Biljana")
	case "biljana":
		fmt.Println("biljana")
		fallthrough
	case "stanoevska":
		fmt.Println("stanoevska")
	case "Stanoevska":
		fmt.Println("Stanoevska")
	default:
		fmt.Println("I don't know")
	}


	// idiomatic way

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got even number", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got odd number", tmpNum)
	default:
		fmt.Println("Number is ", tmpNum)
	}


	// error checking
	_, err := os.Open("./test.txt")

	if err != nil {
		fmt.Println("This is the error code", err)
	}




}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}