package main

import "fmt"

func testingMapDeclarations() {
	// ways to declare a map

	// this is used if the values will be added later
	var colors1 map[string]string

	// this is very similar to ^
	colors2 := make(map[int]string)

	colors2[10] = "#fff"
	colors2[110] = "#fff"
	delete(colors2, 110)
	//!!! cannot do this with a map because we can use numbers as keys
	//! struct.110
	//!!! cannot do this with a map
	fmt.Println(colors1)
	fmt.Println()
	fmt.Println(colors2)
	fmt.Println()
}

func main() {
	// map -> all keys are string, all the values are string
	colors := map[string]string{
		"red":    "#ff0000",
		"white":  "#fff",
		"random": "#565",
	}

	printMap(colors)
}

// will be called with a color map with string keys and string values
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// differences
// a map can have any key
// a map is indexed - a struct is not
// a map is a reference type -- a struct is a value type
// a map does not need to know all the keys that will be on it at compile time

// use a map for any really closely related properties
// use a struct for representing disparate properties held on one location

// structs are generally used more than maps
