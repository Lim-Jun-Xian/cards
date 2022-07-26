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

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver to print the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Multiple return values in one function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Joining a Slice of Strings
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Saving Data to the Hard Drive
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Reading From the Hard Drive
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - Log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffling a Deck using Random Number Generation
func (d deck) shuffle() {
	for i := range d {
		//Generate different Int64 number everytime program start up
		//Use that as a seed to generate a new source object
		source := rand.NewSource(time.Now().UnixNano())

		r := rand.New(source) //Use source object for new random number generator
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
