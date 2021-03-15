package main

import "fmt"

// what functions and return types does a bot have?
type bot interface {
	// any types that have this function, they are now also of type bot
	// because you are of type bot, you can call the functions that are attached to bot
	getGreeting() string
	// more complicated interface type
	// getGreeting(string, int) (string, error)
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// interface function on type bot
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// dont need the 2 values because not using "This"
func (englishBot) getGreeting() string {
	// VERY CUSTOM LOGIC
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
