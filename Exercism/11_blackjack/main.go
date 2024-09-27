package main

import "fmt"

func ParseCard(card string) int {
	switch {
	case card == "ace":
		return 11
	case card == "one":
		return 1
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
// Stand (S)
//Hit (H)
//Split (P)
//Automatically win (W)

func FirstTurn(card1, card2, dealerCard string) string {
	// read the card
	c1 := ParseCard(card1)
	c2 := ParseCard(card2)
	c3 := ParseCard(dealerCard)

	fmt.Println(c1, c2, c3)

	switch {
	// If you have a pair of aces you must always split them.
	case c1 == 11 && c2 == 11:
		println("P")
		return "P"
	case c1+c2 == 21 && c3 == 11 || c1+c2 == 21 && c3 == 10:
		println("S")
		return "S"
	// If you have a Blackjack and the dealer does not have an ace, a figure or a ten then you automatically win
	case c1+c2 == 21 && 11 != c3 || c1+c2 == 21 && 10 != c3:
		println("W")
		return "W"
	// If your cards sum up to a value within the range [17, 20] you should always stand.
	// If your cards sum up to a value within the range [12, 16] you should always stand
	// unless the dealer has a 7 or higher, in which case you should always hit.
	case c1+c2 >= 17 && c1+c2 <= 20 || c1+c2 >= 12 && c1+c2 <= 16 && !(c3 > 6):
		println("S")
		return "S"
	default:
		println("H")
		return "H"
	}
}

func main() {
	//value := ParseCard("ace")
	//fmt.Println(value)
	FirstTurn("ace", "ace", "jack")
	FirstTurn("ace", "king", "ace")
	FirstTurn("five", "queen", "ace")
	FirstTurn("ace", "jack", "ace")
	FirstTurn("ten", "six", "six")
}
