package main

import "fmt"

func main() {
	cards := newDeck()

	handSize := 5
	hand, _ := deal(cards, handSize)

	err := hand.saveToDisk("deck.txt")
	if err != nil {
		fmt.Println("Error: unable to write to file", err)
	}
	shuffledHand := hand.shuffle()
	shuffledHand.print()
}
