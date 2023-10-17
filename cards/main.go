package main

import "fmt"

func main() {
	// cards := newDeck()
	// cards.print()
	// fmt.Println(cards.toString())
	// fmt.Println(cards.deal(4))
	// cards.saveToFile("my_cards")
	cards := loadFromFile("my_cards")
	// cards.print()
	cards.shuffle()
	cards.print()

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range nums {

		if num%2 == 0 {
			fmt.Println("even", num)
		} else {
			fmt.Println("odd", num)
		}
	}

}
