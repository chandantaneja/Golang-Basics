package main

import (
	"fmt"
	"strconv"
)

// Define struct - person
type Person struct {
	firstName, lastName, github, gender  string
	age       int
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Chandan", lastName: "Taneja", github: "chandantaneja", gender: "Make", age: 20}
	fmt.Println(person1)

	// another way of init
	person2 := Person{"Traversy", "Media", "traversymedia", "f", 1}
	fmt.Println(person2)

	// can change the order
	person3 := Person{firstName: "Chandan", github: "chandantaneja", gender: "Male", age: 20, lastName: "Taneja"}
	fmt.Println(person3)

	// can miss any field
	person4 := Person{firstName: "Chandan", github: "chandantaneja", gender: "Male", age: 20}
	fmt.Println(person4)

	// To access the individual fields
	fmt.Println(person1.age)
	fmt.Println(person4.lastName) // The o/p is blank because we havent given any i/p to this field yet
	person4.lastName = "Added"
	fmt.Println(person4.lastName)

	fmt.Println(person4.greet())
	person4.hasBirthday()
	fmt.Println(person4.greet())

	person2.getMarried("LastName")
	fmt.Println(person2.greet())

}

// Greeting method (value reciever)
func (person Person) greet() string {
	return "Hello, my name is " + person.firstName + " " + person.lastName + ". I am " + strconv.Itoa(person.age)
}

// hasBirthday method (pointer receiver)
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "male" {
		return
	} else {
		person.lastName = spouseLastName
	}
}