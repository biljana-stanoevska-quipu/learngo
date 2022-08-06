package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	// anonymous function
	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	// anonymous function
	go func() {
		defer waitGrp.Done()
		fmt.Println("there")
	}()
	waitGrp.Wait()

}