package main

import "fmt"

func divide(a, b float32) float32 {
	return a / b
}

//10/0=infinite
//first error handling for infinte
func divide1(a, b float32) float32 {
	if b == 0 {
		return 0
	}
	return a / b
}

//10/0=infinite
//first error handling for infinte
func divide1(a, b float32) float32 {
	if b == 0 {
		return 0
	}
	return a / b
}

/*
In Go, the underscore(_) is used as a blank identifier. It serves a a write-only variable
that allows you to discard values returned by a function or to ignore specefic values
when you are interested in using them.
*/
func main() {
	fmt.Println("started error Handling in the functions")
	//	ans := divide(10, 4) // result = 2.5
	//	fmt.Println("divide value is : ", ans)
	ans1 := divide1(10, 0)

	fmt.Println("divide value is : ", ans1)
}
