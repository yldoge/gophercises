package main

import (
	"fmt"

	"github.com/yldoge/gophercises/blackjack_ai/blackjack"
)

func main() {
	game := blackjack.New(blackjack.Options{
		Decks:           3,
		Hands:           2,
		BlackjackPayout: 1.5,
	})
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println("Final winnings: ", winnings)
}
