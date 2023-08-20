package game

import "fmt"

type card struct {
	suit  string
	value int
}

func newCard(suit string, value int) *card {
	return &card{suit: suit, value: value}
}

func (c card) getName() string {
	return fmt.Sprintf("%s %d", c.suit, c.value)
}
