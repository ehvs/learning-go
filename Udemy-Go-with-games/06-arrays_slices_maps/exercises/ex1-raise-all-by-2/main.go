/*
Write a program which reads a set of integers,
on a single line, stores the numbers in a slice & raises all elements’ power by 2.

You are given N which represents the count of numbers which come after it.
Example:
1 2 3 4 5

1 4 9 16 25
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin) // adicionar um reader
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	elementsStr := strings.Split(strings.TrimSpace(line), " ")
	var elements []int

	for _, value := range elementsStr { // converter o slice de String para um slice Int
		num, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		elements = append(elements, num)
	}

	for i := 0; i < len(elements); i++ {
		elements[i] *= elements[i]
	}
	for _, value := range elements {
		fmt.Print(value, " ")
	}
}
