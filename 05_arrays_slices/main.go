package main

import "fmt"

func main (){
	// declare
	var fruitArr [2] string

	// assign
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])

	// declare and assign 
	vegArr := [2] string{"Tomato", "Potato"}
	fmt.Println(vegArr)

	// Slices 

	fruitSlice := [] string{"Apple", "Orange", "Grapes", "Cherry"}
	fmt.Println(fruitSlice)

	// To calculate length
	fmt.Println(len(fruitArr))

	// to find the range 
	fmt.Println(fruitSlice[1:2]) // starts at 1 but stops before 2
	fmt.Println(fruitSlice[1:3])

	









}