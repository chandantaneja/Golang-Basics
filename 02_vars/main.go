package main

import "fmt"

func main() {
	/*
		MAIN TYPES
		string
		bool
		int int8 int16 int32 int64
		unit unit8 unit16 unit32 unit64 unitptr
		byte - alias for unit8
		rune - alias for int32
		float32 float64
		complex64 complex128
	*/

	// Creating a variable
	//Using var

	var name string = "Chandan Taneja" // This is a gloal way and can work in any scope (global or local)

	var str = "This is also a String"

	fmt.Println(name)
	fmt.Println(str)

	// we can declare vars even in the middle
	var age int = 20

	fmt.Println(age)

	// String concatenation
	fmt.Println("My name is ", name, ". I am ", age, " years old.")

	// To know the data type of variable
	fmt.Printf("%T \n", name)
	fmt.Printf("%T \n", age)

	var a int32 = 123

	fmt.Printf("%b \n", a)

	// We cannot redefine a variable declared with a keyword const

	const isCool = true

	fmt.Println(isCool)
	// isCool = false

	// Second method of variable declration
	githubUsername := "chandantaneja" // This method of variable declaration can be used only inside a funtion
	fmt.Printf("%T \n", githubUsername)
	fmt.Println(githubUsername)

	size := 1.3
	fmt.Printf("%T\n", size)

	myName, myPortfolio := "Chandan Taneja", "devchandantaneja.co"

	fmt.Println("Name:- ", myName, " Portfolio:- ", myPortfolio)
}
