package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
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

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	fmt.Println("hand: ", d[:handSize])
	fmt.Println("Remaining deck: ", d[handSize:])
	return d[:handSize], d[handSize:]
}

// receiver
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		println("error while reading the file: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for idx := range d {
		//generate a random number
		newIdx := r.Intn(len(d) - 1)
		d[idx], d[newIdx] = d[newIdx], d[idx]
	}
}
