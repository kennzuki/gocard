package main

//class extend slice like oop
type deck []string 

//create or return a new cards
func newDeck() deck {
	cards:=deck{}

	//create cards slice of strings

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	//set the deck
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}


//receiver in a function more like this
func (d deck) print()  {
	for i, card := range d {
		println(i, card)
	}
}
//mulriple return types of dec
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}