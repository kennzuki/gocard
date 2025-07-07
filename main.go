package main



func main() {
cards := deck{
	"Six of Diamonds",
	newCard(),
}
cards=append(cards, "Seven of Diamonds")

cards.print()
}

func newCard() string{
	return "Ace of Spades"
}

