package main

import (
	"fmt"
	"slices"
)

func PreparationTime(layers []string, minutes int) int {
	/*Implement a function PreparationTime that accepts a slice of layers as a []string and
	the average preparation time per layer in minutes as an int.
	The function should return the estimate for the total preparation time based on the number of layers as an int.
	If the average preparation time is passed as 0 (the default initial value for an int),
	then the default value of 2 should be used.*/
	count := len(layers)
	if minutes != 0 {
		println(count * minutes)
		return count * minutes
	} else {
		return count * 2
	}
}

func Quantities(layers []string) (int, float64) {
	/*For each noodle layer in your lasagna, you will need 50 grams of noodles.
	For each sauce layer in your lasagna, you will need 0.2 liters of sauce.
	The function will then determine the quantity of noodles and sauce needed to make your meal.
	The result should be returned as two values of noodles as an int and sauce as a float64.*/
	noodlesCounter := 0
	sauceCounter := 0
	for _, value := range layers {
		if value == "noodles" {
			noodlesCounter++
		}
		if value == "sauce" {
			sauceCounter++
		}
	}
	amountNoodles := noodlesCounter * 50
	amountSauce := float64(sauceCounter) * 0.2
	fmt.Printf("%d, %.1f ", amountNoodles, amountSauce)
	return amountNoodles, amountSauce
}

func AddSecretIngredient(friendsList, myList []string) {
	/* Write a function AddSecretIngredient that accepts two slices of ingredients of type []string as parameters.
	The first parameter is the list your friend sent you, the second is the ingredient list of your own recipe.
	The last element in your ingredient list is always "?". The function should replace it with the last item from your friends list.
	Note: AddSecretIngredient does not return anything -
	you should modify the list of your ingredients directly.
	The list with your friend's ingredients should not be modified.
	Also, since slice is passed into a function as pointers,
	changes to the two []string arguments passed into AddSecretIngredient will be modified directly.*/
	//myList[len(myList)-1] = friendsList[len(friendsList)-1]
	//(*myList)[len(*myList)-1] = (*friendsList)[len(*friendsList)-1]
	// Dereferencing: Use (*myList) and (*friendsList) to dereference the pointers and access the actual slice.
	// OR
	//
	myList = slices.Replace(myList, len(myList)-1, len(myList), friendsList[len(friendsList)-1])
	// i => start index => 5-1= 4
	// j => end index => 5
	//  the range is defined by two indices: a start index (inclusive) and an end index (exclusive).
	// The upper bound is this end index.
	fmt.Println(myList)
}

func ScaleRecipe(recipe []float64, portions int) []float64 {
	/*Implement a function ScaleRecipe that takes two parameters.
	A slice of float64 amounts needed for 2 portions.
	The number of portions you want to cook.
	The function should return a slice of float64 of the amounts needed for the desired number of portions.
	You want to keep the original recipe though.
	This means the quantities argument should not be modified in this function.
	*/
	var newRecipe []float64
	for i := 0; i < len(recipe); i++ {
		newRecipe = append(newRecipe, recipe[i]*float64(portions)/2)
	}
	fmt.Println(newRecipe)
	return newRecipe
}

func main() {
	//layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	//PreparationTime(layers, 3)
	//Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})
	//friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	//myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	//AddSecretIngredient(friendsList, myList)
	quantities := []float64{1.2, 3.6, 10.5}
	ScaleRecipe(quantities, 4)
}
