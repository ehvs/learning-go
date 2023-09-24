package main

import (
	"fmt"
)

func main() {
	//for i := 97; i <= 122; i++ {
	//	fmt.Println(string(rune(i)))
	//}
	// OR
	//for i := rune(97); i <= rune(122); i++ {
	//	fmt.Println(string(i))
	//}
	// OR
	for i := 'a'; i <= 'z'; i++ {
		fmt.Println(string(i))
	}
}
