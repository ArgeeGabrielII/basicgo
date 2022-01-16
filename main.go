package main

func main() {
	cards := newDeck()
	cards.saveToFile("decks_list")
}
