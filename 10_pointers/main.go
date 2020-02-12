package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a // & is used to refer to the address

	fmt.Println(a, b)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b) // the * in the o/p suggests pointer

	// Use * to read val from  address
	fmt.Println(*b)

	// Change val using the pointer
	*b = 10
	fmt.Println(a)

}
