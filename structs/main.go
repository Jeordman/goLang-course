package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// new custom type PERSON that is a struct
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// struct representing a person
	//! go interprets which field by the order
	jeordin0 := person{"Jeordin", "Callister", contactInfo{"@gmail", 5}}
	fmt.Println(jeordin0)
	fmt.Println("")
	// or (this is broken)
	// jeordin1 := person{firstName: "Jeordin", lastName: "Callister", contactInfo{email: "@gmail.com", zipCode: 840000}}

	// or -- this assigns a zero value to each field
	var rylie person
	rylie.firstName = "Rylie"
	rylie.lastName = "Fugal"
	// show each field and their value
	fmt.Printf("%+v", rylie)
	fmt.Println("")

	jeordin := person{
		firstName: "Jeordin",
		lastName:  "Callister",
		contactInfo: contactInfo{
			email:   "@gmail.com",
			zipCode: 64,
		},
	}

	// the & is an operator
	// means -> give me the memory address this variable is pointing at
	jeordinPointer := &jeordin
	jeordinPointer.updateName("Jeo")
	jeordin.print()

	fmt.Println("")

	// there's is an easier way to write this ^
	// this is the shorthand
	jeordin.updateName("Test")
	jeordin.print()

}

// * is an operator
// give me the value that exists at this memory adress
// type is changed to -> pointer pointing at a person || you can just put in a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// this does not actually update jeordin
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
