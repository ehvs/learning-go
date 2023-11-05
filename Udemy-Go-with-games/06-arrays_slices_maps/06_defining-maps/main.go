package main

import "fmt"

func main() {
	// this a slice
	playerPoints := []int{
		10,
		50,
		30,
	}

	fmt.Println(playerPoints)

	// como inicializar um map
	myMap := map[string]int{}
	myMap["my age"] = 31
	fmt.Println(myMap)

	// this is a map
	// Syntax: this map has keys of type 'string' and values of type 'int'
	playerPointsmap := map[string]int{
		"player1": 10,
		"player2": 50,
		"player3": 30,
	}
	fmt.Println("Map:\n", playerPointsmap)

	p1points := playerPointsmap["player1"]
	fmt.Println("Player 1 has:", p1points, "points.")

	p4points := playerPointsmap["player4"]            // player does not exist
	fmt.Println("Player 4 has:", p4points, "points.") // returns 0 by default

	p4points, ok := playerPointsmap["player4"]
	fmt.Println(p4points, ok) // Printing a bool, returns 0 as it does not exist.
	//if !ok {                  // Se nao estiver OK, entao panic.
	//	panic("this user does not exist")
	//}

	_, ok = playerPointsmap["player4"]
	if !ok {
		playerPointsmap["player4"] = 50 // adding a value to map
		fmt.Println("Added player 4", playerPointsmap)

	}

	// OR

	if _, ok = playerPointsmap["player5"]; !ok {
		playerPointsmap["player5"] = 12
		fmt.Println("Added player 5", playerPointsmap)
	}

	// Deleting
	player, ok := playerPointsmap["player1"]
	fmt.Println(player, ok)
	delete(playerPointsmap, "player1")
	fmt.Println("Map:\n", playerPointsmap) // demonstra que Player 1 foi deletado

	player, ok = playerPointsmap["player1"]
	fmt.Println(player, ok) // demonstra que a existencia de Player 1 eh FALSA
}
