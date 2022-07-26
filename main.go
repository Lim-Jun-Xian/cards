package main

func main() {
	/*
		cards := deck{"Ace of Diamonds", newCard()} //create variable with a slice of type string
		cards = append(cards, "Six of Spades")      // Add the value Six of Spades into the cards variable
	*/

	/* Multiple return values in one function
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
	*/

	/* Joining a Slice of Strings
	cards := newDeck()
	fmt.Println(cards.toString())
	*/

	/* Saving Data to the Hard Drive
	cards := newDeck()
	cards.saveToFile("my_cards")
	*/

	/* Reading From the Hard Drive
	cards := newDeckFromFile("my")
	cards.print()
	*/
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

/*
// Define a function called 'newCard', when executed, this function will return a value of type 'string'
func newCard() string {
	return "Five of Diamonds"
}
*/
