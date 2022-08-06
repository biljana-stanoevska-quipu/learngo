package main

import (
	"fmt"
	"time"
)

func main() {

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("DONE!")
			break
		}
		fmt.Println(timer)
		time.Sleep(time.Second)
	}

	friends_out_mkd := []string{
		"ivana",
		"angela",
		"ninka",
		"tina"}

	friends_in_mkd := []string{
		"ninka",
		"tina"}

	breakPoint:
		for index, value := range(friends_out_mkd) {
			fmt.Println(index, value)
			for _, value_new := range(friends_in_mkd){
				if value == value_new {
					fmt.Println("Not possible.", value_new, "is out and in")
					break breakPoint
				}
			}
		}

}