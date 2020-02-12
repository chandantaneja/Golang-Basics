package main

import "fmt"

func main(){
	ids := [] int{33, 76, 15, 200, 45, 65, 100}

	for i, id := range ids{
		fmt.Printf("%d - ID: %d \n", i, id)
	}

	for _, id := range ids{
		fmt.Printf("ID: %d \n", id)
	}
	
	// Add all the ids
	sum := 0
	for _, id :=range ids{
		sum += id
	}

	fmt.Println("Sum", sum)

	// range with map 
	// This is also a way to declare map

	emails := map[string]string{"Bob": "bob@gmail.com", "Ron":"Ron@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}