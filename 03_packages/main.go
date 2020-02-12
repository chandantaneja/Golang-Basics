package main

//Correct way of putting packages

import (
	"fmt"
	"math"
	"github.com/chandantaneja/go_crash_course/03_packages/strutil"
)

func main(){
	fmt.Println("Hello World !!!")
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(4))
	fmt.Println(strutil.Reverse("dlrow olleh"))
	fmt.Println(strutil.Reverse("hello World"))
}