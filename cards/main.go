package main

func main() {
	cards := newDeck()

	handSize := 5
	hand, _ := cards.deal(handSize)

	err := hand.saveToDisk("deck.txt")
	if err != nil {
		panic("unable to write to file")
	}
	shuffledHand := hand.shuffle()
	shuffledHand.print()
}
