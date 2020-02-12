package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// creates a Buffer Reader Object

func main() {
	var myString string
	fmt.Scanln(&myString)
	fmt.Println(myString)

	var name string = "Chandan Taneja"
	var a, _ = fmt.Println(name)
	fmt.Println(a)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Full name : ")
	myName, _ := reader.ReadString('\n')
	fmt.Println(myName)

	fmt.Print("Enter your rating: ")
	myRating, _ := (reader.ReadString('\n'))
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(myRating), 64)
	fmt.Println(myNumRating + 2) 
}
