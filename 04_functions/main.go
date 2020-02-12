package main

import "fmt"

func greeting(name string) string{
	return "Hello " + name 

}

func main(){
	fmt.Println(greeting("Chandan"))
	fmt.Println(getSum(5,10))
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}