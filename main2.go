package main

func main() {
	cards := newDeck()
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "six of spades")

	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}

	hand, remainingcards := deal(cards, 5)
	hand.print()
	remainingcards.print()
}
