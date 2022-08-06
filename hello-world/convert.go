package main

import(
	"fmt"
	"strings"
)

func main() {
	author := "Brayan Adams"
	course := "Go"

	fmt.Println(converter(author, course))

	fmt.Println(author)
	fmt.Println(course)

	bestFinish := championshipFinishes(12, 6, 5, 4, 3, 3)
	fmt.Println(bestFinish)

	if bestFinish < 4 {
		fmt.Println("Vert fast.")
	} else {
		fmt.Println("No so fast.")
	}

}

func converter(author, course string) (a, c string) {

	author = strings.ToUpper(author)
	course = strings.Title(course)
	return author, course
}

func championshipFinishes(finishes ...int) int {

	best := finishes[0]

	for _, i := range(finishes) {
		if i < best {
			best = i
		}
	}

	return best

}