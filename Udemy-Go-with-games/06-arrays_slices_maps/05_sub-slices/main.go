package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	// cortando os tres primeiros index
	subslice := numbers[0:2] // or [:2]

	fmt.Println(subslice)

	// cortando a partir do index 2 ate o final
	subsliceEnd := numbers[2:len(numbers)] // or [2:]
	fmt.Println(subsliceEnd)
}

// essa syntaxe possibilita uma "view" do slice original
// ou seja, se modificar um index na "view", sera tb modificado na orignal.
