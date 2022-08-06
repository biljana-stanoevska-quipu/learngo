package main

import (
	"fmt"
)

func main() {

	friendsCountries := make(map[string]string)

	friendsCountries["ivana"] = "switzerland"
	friendsCountries["angela"] = "austria"
	friendsCountries["ninka"] = "usa"


	friendsCountries2 := map[string]string{
		"jasna": "croatia",
		"viki": "australia"}

	fmt.Printf("friendsCountries: %v\nfriendsCountries2: %v\n", friendsCountries, friendsCountries2)

	for mapKey, mapVal := range friendsCountries {
		fmt.Printf("Key is: %v. Value is: %v\n", mapKey, mapVal)
	}

	fmt.Println(friendsCountries2["jasna"])

	friendsCountries2["jasna"] = "mkd"
	fmt.Println(friendsCountries2["jasna"])

	friendsCountries2["maja"] = "germany"
	fmt.Println(friendsCountries2)

	delete(friendsCountries2, "jasna")
	fmt.Println(friendsCountries2)


}