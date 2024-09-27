package main

import (
	"fmt"
	"slices"
)

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	favCards := []int{2, 6, 9}
	fmt.Printf("Array is %d", favCards)
	return favCards
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index >= 0 && index <= len(slice) {
		n := slice[index]
		fmt.Printf("\ncard is: %d", n)
		return n
	} else {
		return -1
	}
}

func SetItem(slice []int, index, value int) []int {
	if index >= 0 && index < len(slice) {
		slice[index] = value
		fmt.Printf("\nnew value is: %d", slice[index])
		return slice
	} else {
		slice = append(slice, value)
		return slice
	}
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	slice = slices.Insert(slice, 0, values...)
	return slice
}

func RemoveItem(slice []int, index int) []int {
	if index >= 0 && index <= len(slice) {
		slice = slices.Delete(slice, index, index+1)
		return slice
	} else {
		return slice
	}
}

func main() {
	//FavoriteCards()
	//GetItem([]int{1, 2, 4, 1}, 1) // card == 4
	//index := 4
	//newCard := 6
	//cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
	//fmt.Println(cards)
	//slice := []int{3, 2, 6, 4, 8}
	//cards := PrependItems(slice, 5, 1)
	//fmt.Println(cards)
	cards := RemoveItem([]int{3, 2, 6, 4, 8}, 7)
	fmt.Println(cards)
}
