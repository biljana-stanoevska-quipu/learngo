package main

import (
    "fmt"
    "reflect"
    "strconv"
    "os"
)

func main() {

    name := os.Getenv("USERNAME")  // "Biljana"
    course := "Go fundamentals"
    module := "4"  // line comment
    clip := 2
    // courseComplete := false

    fmt.Println("name and course are set to", name, "and", course, ".")
    fmt.Println("module and clie are set to", module, "and", clip, ".")
    fmt.Println("name is of type", reflect.TypeOf(name))
    fmt.Println("module is of type", reflect.TypeOf(module))

    // total := module + clip
    // fmt.Println("module + clip equals", total)

    iModule, err := strconv.Atoi(module)  // Atoi is ASCII string to integer
    if err == nil {
        total := iModule + clip
        fmt.Println("module + clip equals", total)
    }

    fmt.Println("Memory address of *course* variable is", &course)

    var ptr *string = &course  // * says that ptr varibale is a pointer to string, ptr is a memory address
    fmt.Println("Memory address of ptr is", ptr, "Value stored is", *ptr)  // with * dereference


    fmt.Println("\nHi", name, "your current course is", course)
    // updateCourse(course)  // passes a copy of course variable
    updateCourse(&course)  // passes by reference
    fmt.Println("\nHi again", name, "your current course is", course)

}

func updateCourse(course *string) string {
    *course = "Getting started with Docker"
    fmt.Println("Updated course to", *course)
    return *course
}