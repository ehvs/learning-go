package main

import (
	"fmt"
	"reflect"
	"strconv"
)

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

	afloat := 3.5
	var anotherInt int = int(afloat)
	fmt.Println(anotherInt) // result is 3

	// converting string to int
	var numAsString string = "50"
	numInt, err := strconv.Atoi(numAsString) //Atoi returns the result + error so both must be declared and 'collected'
	fmt.Println("Becoming a int:", numInt, err)
	fmt.Println("The proof!", numInt, "is:", reflect.ValueOf(numInt).Kind())

	// converting int to string
	var numberAsString string = strconv.Itoa(numInt)
	fmt.Println("Becoming a integer again:", numberAsString)
	fmt.Println("The proof!", numberAsString, "is:", reflect.ValueOf(numberAsString).Kind())
}
