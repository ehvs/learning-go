package main

import (
	"fmt"
)

func sumDigits(number int) int {
	res := 0
	for number > 0 {
		res += number % 10
		fmt.Println("res:", res)
		number /= 10
		fmt.Println("number:", number)
	}
	return res
}

func main() {
	fmt.Println(sumDigits(123))
}
