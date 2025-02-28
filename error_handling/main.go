package main

import "fmt"

func divide(a, b float32) float32 {
	return a / b
}

// 10/0=infinite
// first error handling for infinte
func divide1(a, b float32) float32 {
	if b == 0 {
		return 0
	}
	return a / b
}

// 10/0=infinite
// first error handling for infinte
func divide2(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

//Error handling with String
func divide3(a, b float32) (float32, string) {
	if b == 0 {
		return 0, "denominator must not be zero"
	}
	return a / b, "nil"
}

/*
In Go, the underscore(_) is used as a blank identifier. It serves a a write-only variable
that allows you to discard values returned by a function or to ignore specefic values
when you are interested in using them.
*/
func main() {
	fmt.Println("Started Error Handling in the Functions")
	//	ans := divide(10, 4) // result = 2.5
	//	fmt.Println("divide value is : ", ans)

	ans0 := divide(10, 0)
	ans1 := divide1(10, 0)

	fmt.Println("divide value is for ans0; ", ans0) // divide value is;  +Inf
	fmt.Println("divide value is for ans1: ", ans1) //divide value is :  0
	

	ans3, error := divide2(10, 0)
	if error != nil {
		fmt.Println("Error handling")
	}
	fmt.Println("divide value is for ans3: ", ans3)

	//if i dont need to handle error then we can use _ (underscore)
	// ans3,_ := divide2(10,0)
	// fmt.Println("divided value is for ans ans3: ", ans3)

	//Another way to handle error with String

	ans4, string := divide3(10,0);
	if(string != ""){
		fmt.Println("error handling for string ")

	}
	fmt.Println("divided value is for : ", ans4)
}
