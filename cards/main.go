// In vim run :GoTest to get warnings
// GoLint --> lint project
// GO IS NOT OBJECT ORIENTED
package main

//old -----
// create var - print it out
// var testCard string = "Ace of Spades"

// infer type
// Only use this := when creating a variable
// card := newCard()

// fmt.Println(card)
// -----

func main() {
	// slice of type deck
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	cards.saveToFile("my_cards")
}
