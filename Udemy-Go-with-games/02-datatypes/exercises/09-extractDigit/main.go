package main

import (
	"fmt"
)

func main() {
	var number int = 123
	// use modulus operator to extract digit
	// % modulo-divides 2 varilables : 1e2
	// 1 digit : 1e1
	lastDigit := number % 1e1
	fmt.Println("(Given 123)")
	fmt.Println(lastDigit)
}
