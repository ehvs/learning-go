package main

import "fmt"

func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	fmt.Printf("The sum is: %d\n", result)
}
