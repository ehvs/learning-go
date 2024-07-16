package main

import "fmt"

func main() {
	primes := map[int]bool{ // this map have values of type 'boolean'
		2: true,
		3: false,
		4: false,
		5: true,
	}

	// Iterating with maps
	// In a FOR loop with maps, the 1st variable is the 'key', and the 2nd is the 'key'
	// It is not an ordered list

	for key, value := range primes {
		fmt.Println(key, "->", value)
	}
}
