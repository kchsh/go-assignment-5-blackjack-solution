package game

import (
	"math/rand"
	"time"
)

type deck struct {
	cards []card
}

func newDeck() *deck {
	d := &deck{}
	d.generateDeck()

	return d
}

func (d *deck) generateDeck() {
	var cardSuits []string = []string{"Червы", "Пики", "Бубны", "Крести"}
	var cardValues []string = []string{"Туз", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Джек", "Королева", "Король"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d.cards = append(d.cards, *newCard(suit, getCardValue(value)))
		}
	}
}

func (d *deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	for i := 0; i < len(d.cards); i++ {
		newIndex := rand.Intn(len(d.cards))
		c := d.cards[i]
		d.cards[i] = d.cards[newIndex]
		d.cards[newIndex] = c
	}
}

func (d *deck) dealCard() card {
	topCard := d.cards[0]
	d.cards = d.cards[1:]

	return topCard
}

func getCardValue(cardValue string) int {
	if cardValue == "Туз" {
		return 1
	} else if cardValue == "2" {
		return 2
	} else if cardValue == "3" {
		return 3
	} else if cardValue == "4" {
		return 4
	} else if cardValue == "5" {
		return 5
	} else if cardValue == "6" {
		return 6
	} else if cardValue == "7" {
		return 7
	} else if cardValue == "8" {
		return 8
	} else if cardValue == "9" {
		return 9
	} else if cardValue == "10" {
		return 10
	} else if cardValue == "Джек" {
		return 10
	} else if cardValue == "Королева" {
		return 10
	} else if cardValue == "Король" {
		return 10
	} else {
		return -1
	}
}
