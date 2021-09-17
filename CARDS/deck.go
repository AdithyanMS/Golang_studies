package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Spades", "Clubs", "Diamonds"}
	cardvalues := []string{"Ace", "2", "3", "4"}

	for _, suit := range cardSuits {
		for _, value := range cardvalues {
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
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error is: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newIndex := rand.Intn(len(d) - 1)
		d[i], d[newIndex] = d[newIndex], d[i]

	}
}
