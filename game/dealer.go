package game

import "fmt"

type dealer struct {
	name string
	hand *hand
}

func newDealer() *dealer {
	return &dealer{name: "Компьютер", hand: newHand()}
}

func (p *dealer) isBusted() bool {
	return p.hand.getTotalValue() > 21
}

func (p *dealer) getHandValue() int {
	return p.hand.getTotalValue()
}

func (p *dealer) addCard(c card) {
	if len(p.hand.cards) > 0 {
		fmt.Printf("Дилер [%s] получил карту [%s]\n", p.name, c.getName())
	}

	p.hand.addCard(c)
}
