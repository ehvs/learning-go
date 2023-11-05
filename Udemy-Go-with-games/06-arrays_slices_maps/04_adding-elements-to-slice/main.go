package main

import (
	"fmt"
	"sort"
)

func main() {
	zombies := []string{}
	zombies = append(zombies, "Paul", "George", "Katya")
	zombiesCopy := append(zombies, "Lucy")
	fmt.Println(zombies)
	fmt.Println(zombiesCopy)

	// Combining slices
	odds := []int{1, 3, 5, 7}
	evens := []int{2, 4, 6, 8}
	newSlice := append(odds, evens...)
	// Sorting
	sort.Ints(newSlice)
	fmt.Println(newSlice)
}
