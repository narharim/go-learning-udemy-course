package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func init() {
	rand.Seed(time.Now().UnixMilli())
}
func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Heart", "Spades", "Diamond", "Club"}
	cardValue := []string{"Ace", "Two", "Three", "Four"}

	for _, value := range cardValue {
		for _, suit := range cardSuits {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, value := range d {
		fmt.Println(value)
	}
}

func (d deck) deal(h int) (deck, deck) {
	return d[:h], d[h:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToDisk(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromDisk(filename string) deck {

	deckBytes, err := os.ReadFile(filename)
	if err != nil {
		panic("unable to load the file")
	}
	return strings.Split(string(deckBytes), ",")

}

func swap(d deck) deck {
	for index := range d {
		newIndex := rand.Intn(len(d) - 1)
		d[index], d[newIndex] = d[newIndex], d[index]
	}
	return d
}

func (d deck) shuffle() deck {
	return swap(d)
}
