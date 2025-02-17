package main

import "fmt"

func main() {
	age := 12
	name := "majharul"
	height := 4.444

	fmt.Println("age: ", age, "name: ", name, "height: ", height)

	//Printf(print formatted)
	//format specifiers
	/*
		%d => Integer
		%s => String
		%T => Types of the value
		%.3f => Float values

	*/

	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.7f\n",height) 
	// %.f => print will only 4
	// %.1f => print will after 1 digit : 4.4
	// %.7f => print will after 7 digit : 4.4440000
	fmt.Printf("name is : %s\n", name)

	fmt.Printf("height is %T\n",height) // will return type of value.

	fmt.Println("age: ", "name : ", age, name)
	// age:  name :  12 majharul
	fmt.Printf("age: %d, name : %s", age, name)


}
