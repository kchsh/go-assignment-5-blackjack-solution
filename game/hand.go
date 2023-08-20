package game

type hand struct {
	cards []card
}

func newHand() *hand {
	return &hand{}
}

func (h *hand) addCard(card card) {
	// Если туз и сумма руки будет меньше 21
	if card.value == 1 && h.getTotalValue()+11 <= 21 {
		card = *newCard(card.suit, 11)
	}

	h.cards = append(h.cards, card)
}

func (h hand) getTotalValue() int {
	sum := 0
	for _, card := range h.cards {
		sum += card.value
	}

	return sum
}
