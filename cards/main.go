package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	cards := []string{newCard(), newCard()}
	cards = append(cards, newCard())
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}