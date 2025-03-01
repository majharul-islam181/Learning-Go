/*
	(1) In Go, a slice is a fliexible and dynamic data structure that provides a more
	powwerful alternative to arrays.

	(2) Unlike arrays, slices have a dynamic size,and their length can be changed during
	the program's during the program execution.

*/

/*

	1. Dynamic Size: Slices can grow or shrink dynamically during runtime.
	2.
	3. Syntax: numbers := []int {1,2,3}
	4. Use of Make Function: Can make a function to create a slice with specfic length and capasity.
	5. Append : Used to add elements to a slice.

*/

package main

import "fmt"

func main() {

	/*
		numbers := []int{1, 2, 3, 4, 5}
		numbers = append(numbers, 2, 3, 4, 4, 5)

		fmt.Println("Numbers is : ", numbers)
		//Output:  Numbers is :  [1 2 3 4 5]
		fmt.Printf("Numbers has data type: %T\n", numbers)
		//Output:  Numbers has data type: []int
		fmt.Println("length: ", len(numbers))
		//Output: length:  10

		//if i want to pass empty array
		emptyArray := []int{}
		fmt.Println("Empty array : ", emptyArray)
		//Output: Empty array :  []
	*/

	//Create Slice with MAKE function

	numbers := make([]int, 3, 5)
	/*  	array Type,lenght,capacity */
	fmt.Println("Slice: ", numbers)
	// Output : 	Slice:  [0 0 0]
	fmt.Println("Length: ", len(numbers))
	// Output : Length:  3
	fmt.Println("Capacity: ", cap(numbers))
	// Output: Capacity:  5

	//Now append some value in array
	numbers = append(numbers, 2)
	fmt.Println("After adding value Slice will : ", numbers)
	// Output : After adding value :  [0 0 0 2]
	fmt.Println("Length: ", len(numbers))
	// Output : Length:  3
	fmt.Println("Capacity: ", cap(numbers))
	// Output: Capacity:  5

	//Now append some value in array
	numbers = append(numbers, 10, 20, 30)
	fmt.Println("After adding value Slice will : ", numbers)
	// Output : After adding value Slice will :  [0 0 0 2 10 20 30]
	fmt.Println("Length: ", len(numbers))
	// Output : Length:  7
	fmt.Println("Capacity: ", cap(numbers))
	// Output: Capacity:  10

	/* When we adding some value in array it will double in Capacity */
	/*
		Suppose you declare 5 capacity first, then you put more then 5 element,
		it will automatic upadte then Capacity will 10,
		and again you append more then 10 value then it will increase and Capacity will 20 (always double)
	*/

}
