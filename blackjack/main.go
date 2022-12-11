package main

import (
	"fmt"

	"yldoge.com/deck"
)

func main() {
	cards := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	for i := 0; i < 10; i++ {
		card, cards = cards[i], cards[1:]
		fmt.Println(card)
	}
}
