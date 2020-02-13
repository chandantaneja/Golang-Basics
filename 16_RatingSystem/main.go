package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	// Front End

	// take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name :")
	name, _ = reader.ReadString('\n')

	// take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our center (b/w 1-5)")
	userRating, _ = reader.ReadString('\n')
	myNumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	// Back End
	fmt.Printf("Hello %v \n Thanks for rating our center with %v star rating. \n\n Your rating was recorded in our system at %v\n\n", name, myNumRating, time.Now().Format(time.Stamp))

	if myNumRating == 5 {
		fmt.Println("Excellent")
	} else if myNumRating == 4 || myNumRating == 3 {
		fmt.Println("Good")
	}else if myNumRating < 3{
		fmt.Println("Will do better")
	}
}
