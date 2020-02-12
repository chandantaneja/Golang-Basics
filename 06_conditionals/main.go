package main

import "fmt"

func main(){
	x := 50
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d \n", x, y)
	}else{
	fmt.Printf("%d is less than %d \n", y, x)
	}

	// if else if
	color := "red"

	if color == "red" {
		fmt.Println("color is red")
	}else if color == "blue" {
		fmt.Println("Color is Blue")
	}else {
		fmt.Println("Unknown color")
	}

	color = "red"


	// Switch
	switch color {
	case "red":
	case "blue":
		fmt.Println("color is blue")
		break
	default: 
		fmt.Println("Unknown")
	}
}



