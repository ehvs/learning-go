package main

import (
	"fmt"
)

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int
	for i := 0; i < len(birdsPerDay); i++ {
		totalBirds += birdsPerDay[i]
	}
	fmt.Println(totalBirds)
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// week is 0 to 6 => 7 days
	start := (week - 1) * 7
	end := week * 7

	total := 0
	for _, count := range birdsPerDay[start:end] {
		total += count
	}
	fmt.Println(total)
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		fmt.Print("index is: ", i)
		value := birdsPerDay[i]
		fmt.Println("value: ", value)
		birdsPerDay[i] = value + 1
		fmt.Println("value +1: ", birdsPerDay[i])
	}
	//fmt.Println(birdsPerDay)
	return birdsPerDay
}

func main() {
	//birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	//TotalBirdCount(birdsPerDay)
	//birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	//BirdsInWeek(birdsPerDay, 2)
	birdsPerDay := []int{2, 5, 0, 7, 4, 1}
	FixBirdCountLog(birdsPerDay)
}
