package main

import (
	"blackjack/game"
	"fmt"
)

func main() {
	var input string
	b := game.NewBlackjack()

	for {
		b.StartGame()

		fmt.Println("Сыграем еще раз (д/н)?")
		fmt.Scan(&input)

		if input == "н" {
			break
		}

		b.ClearTable()
		fmt.Println()
	}
}
