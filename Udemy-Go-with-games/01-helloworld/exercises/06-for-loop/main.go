package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")
	time.Sleep(5 * time.Second)
	num := 1000
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}
