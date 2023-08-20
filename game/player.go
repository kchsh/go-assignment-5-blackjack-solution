package game

import "fmt"

type player struct {
	name string
	hand *hand
}

func newPlayer(name string) *player {
	return &player{name: name, hand: newHand()}
}

func (p *player) isBusted() bool {
	return p.hand.getTotalValue() > 21
}

func (p *player) getHandValue() int {
	return p.hand.getTotalValue()
}

func (p *player) addCard(c card) {
	fmt.Printf("Игрок [%s] получил карту [%s]\n", p.name, c.getName())
	p.hand.addCard(c)
}
