package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	for i := 0; i < len(numbers); i++ {
		numbers[i] += 1
	}
	fmt.Println(numbers)

	myslice := []int{2, 3, 4, 5}
	for a, element := range myslice {
		fmt.Println("element: ", element)
		myslice[a] += 1
	}
	fmt.Println(myslice)
}

// element eh uma variavel read-only - nao se modifica
