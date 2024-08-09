package main

import "fmt"

func main() {
	//cards := deck{"Ace of diamonds", "Five of Spades"}
	cards := newDeck()
	cards.print()
	hand, remDeck := deal(cards, 3)
	hand.print()
	remDeck.print()

	fmt.Println("------------")
	fmt.Println(cards.toString())
	cards.saveToFile("umesh")

	diskDeck := newDeckFromFile("umesh")
	diskDeck.print()

	cards.shuffle()
	fmt.Println("-----shuffling cards-------")
	cards.print()
}
