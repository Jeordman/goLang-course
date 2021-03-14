package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{
		"Ace",
		"One",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	// _ = variable not needed
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// create a function that belongs to the deck type
// reciever functions
// any type DECK gets acces to print function
// d --> copy of the deck running the function (this)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// returning 2 values of type deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// return a single string with commas in between each word
	// type conversion []string(var)
	return strings.Join([]string(d), ",")
}

// save deck to hardware
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	// the error object is nil or string
	// bs - short for byte slice
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		// there was an error
		// Option #1 - log the error and return a call to newDec()
		//! CHOSE THIS Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// type convert to string, then split into slice of sentences
	s := strings.Split(string(bs), ",")
	// change slice of strings to deck
	return deck(s)
}

func (d deck) shuffle() {

	// generate a truly randome int64 use it as a seed for a source
	source := rand.NewSource(time.Now().UnixNano())
	// r of type rand is a seeded random generator
	r := rand.New(source)

	for i := range d {
		// this is a sudo random generator (uses a seed value)
		newPositition := r.Intn(len(d) - 1)
		// swap what is held at each postition
		d[i], d[newPositition] = d[newPositition], d[i]
	}
}
