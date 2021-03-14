package main

import "fmt"

// new custom type PERSON that is a struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	// struct representing a person
	//! go interprets which field by the order
	// jeordin := person{"Jeordin", "Callister"}
	// or
	jeordin := person{firstName: "Jeordin", lastName: "Callister"}

	fmt.Println(jeordin)
}
