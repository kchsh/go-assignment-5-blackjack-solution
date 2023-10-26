package game

import "fmt"

type blackjack struct {
	deck   *deck
	player *player
	dealer *dealer
}

func NewBlackjack() *blackjack {
	return &blackjack{deck: newDeck()}
}

func (b *blackjack) Play() {
	b.deck.shuffle()

	player := newPlayer("Костя")
	player.addCard(b.deck.dealCard())
	player.addCard(b.deck.dealCard())

	dealer := newDealer()
	dealer.addCard(b.deck.dealCard())
	dealer.addCard(b.deck.dealCard())

	b.player = player
	b.dealer = dealer

	for {
		fmt.Println("Взять еще одну карту (д/н)?")
		var input string
		fmt.Scan(&input)

		if input == "н" {
			break
		}

		player.addCard(b.deck.dealCard())
		if player.isBusted() {
			break
		}
	}

	if player.isBusted() {
		b.endGame()
		return
	}

	for dealer.getHandValue() <= 17 {
		dealer.addCard(b.deck.dealCard())
	}

	if dealer.isBusted() {
		b.endGame()
		return
	}

	b.endGame()
}

func (b blackjack) endGame() {
	fmt.Printf("\nСумма руки игрока [%s] - %d. Сумма руки дилера - %d\n", b.player.name, b.player.getHandValue(), b.dealer.getHandValue())

	if b.player.isBusted() {
		fmt.Printf("Игрок [%s] сгорел! Дилер [%s] выигрывает.\n", b.player.name, b.dealer.name)
		return
	}
	if b.dealer.isBusted() {
		fmt.Printf("Дилер [%s] сгорел! Игрок [%s] выигрывает.\n", b.dealer.name, b.player.name)
		return
	}

	if b.player.getHandValue() > b.dealer.getHandValue() {
		fmt.Printf("Игрок [%s] выиграл!\n", b.player.name)
	} else {
		fmt.Printf("Игрок [%s] проиграл!\n", b.player.name)
	}
}
