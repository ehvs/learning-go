package main

import "fmt"

func main() {
	zombies := []string{"Paul", "Goeey", "George", "Lucy"} // Iniciando um slice onde os elementos sao do tipo string
	fmt.Println(zombies)
	paul := zombies[0]
	george := zombies[2]
	fmt.Println(zombies, paul, george)

	index := 2
	george = zombies[index]
	fmt.Println(george)

	zombies[1] = "Katya"
	fmt.Println(zombies)
	fmt.Println(len(zombies), zombies)
}

// caveats:
// slices sao "tipados" , ou seja so recebem do tipo que sao inicializados
// sejam string ou int
