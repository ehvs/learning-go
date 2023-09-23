package main

import "fmt"

func main() {
	var intNum int8 = 42
	var biggerInt int16 = int16(intNum)
	fmt.Println(biggerInt)

	//var message string = "Hello"
	//var aNum int = int(message)
	//fmt.Println(aNum)

	var bigNum int16 = 300
	var smallNum int8 = int8(bigNum)
	fmt.Println(smallNum) // result is 44 because the int16 is bigger than int8

	var anint int16 = 42
	var aFloat float32 = float32(anint)
	fmt.Println(aFloat)
}
