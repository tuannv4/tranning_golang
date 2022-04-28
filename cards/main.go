package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards := newDeckFromFile("my_cards1")
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
