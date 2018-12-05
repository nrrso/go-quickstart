package main

import (
	"fmt"
)

func main() {
	// Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and Assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)

	fruitSlice := []string{"Apple", "Orange", "Grape", "Mango"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
