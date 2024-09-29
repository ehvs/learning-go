package main

import "fmt"

type File []bool // is a slice that will contain 8 bool values
type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	counter := 0
	if validFile, exists := cb[file]; exists {
		for _, file := range validFile {
			if file {
				counter++
			}
		}
		fmt.Println("Counter is", counter)
		return counter
	}
	return 0
}

func CountInRank(cb Chessboard, rank int) int {
	counter := 0
	if rank >= 1 && rank <= 8 {
		for _, value := range cb {
			fmt.Println("BEFORE: Value is", value)
			if value[rank-1] {
				counter++
				fmt.Println("AFTER: Value is", value)
			}
		}
		fmt.Println("Counter is", counter)
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	fmt.Println("all is: ", len(cb)*len(cb))
	return len(cb) * len(cb)
}

func CountOccupied(cb Chessboard) int {
	// It should count how many squares are occupied in the chessboard.
	// Return an integer.
	counter := 0
	for _, file := range cb {
		fmt.Println("File is:", file)
		for _, value := range file {
			fmt.Println("value is:", value)
			if value {
				counter++
			}
		}
	}
	fmt.Println("Counter is", counter)
	return counter
}

func forFun(cb Chessboard) {
	for key, value := range cb {
		fmt.Println("Key:", key, "Value:", value)
	}

	// Iterating over the map while ignoring the key
	for _, value := range cb {
		fmt.Println("Value:", value)
	}
}

func main() {
	board := Chessboard{
		"A": make(File, 8),
		"B": make(File, 8),
		"C": make(File, 8)}
	board["A"][0] = true // Occupied
	board["A"][1] = true // Occupied
	board["A"][2] = true // Occupied
	board["B"][3] = true // Occupied
	board["B"][4] = true // Occupied
	board["C"][2] = true // Occupied
	CountInFile(board, "A")
	CountInRank(board, 1)
	CountAll(board)
	CountOccupied(board)
	forFun(board)
}
