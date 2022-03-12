package main

func main() {

	// you can declare varables in two ways:
	// 1- var card string = "Ace of Spades"

	// only works when the declaring a variable not reassigning it
	// make the go compiler to infer the type of the variable
	// 2- :
	// card := "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)

	// create a slice of type strings
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// fmt.Println(cards)
	// print the slice
	// for i, card := range cards{
	// 	fmt.Println(i,card)
	// }

	// cards := newDeck()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	cards := newDeck()
	// fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

}

// return value of type string
// func newCard() string {
// 	return "Five of Diamonds"

// }
