package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Ron"] = "Ron@gmail.com"
	emails["Duke"] = "Duke@yahoo.com"
	emails["Bucky"] = "bucky@outlook.com"
	emails["Chandan"] = "Chandan@golang.com"

	// automatically arranged in Alphabetical order
	fmt.Println(emails)

	//length
	fmt.Println(len(emails))

	fmt.Println(emails["Chandan"])

	// Deleting
	delete(emails, "Bob")
	fmt.Println(emails)
}
