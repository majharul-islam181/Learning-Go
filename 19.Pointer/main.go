package main

import (
	"fmt"
)

/*
In Go, a pointer is a variable that stores the memory address of another variable.
They provide a way to work with the memory directly.
*/

func modifyValueByReference(num *int) {
	*num = *num * 20
}

func main() {
	var num int
	num = 2
	// fmt.Println(num)

	// var ptr *int
	ptr := &num
	fmt.Println("Num has value : ", num)
	fmt.Println("pointer contains : ", ptr)

	//Alternative writting
	num1 := 2
	ptr1 := &num1
	fmt.Println("Num has value : ", num1)
	fmt.Println("pointer contains : ", ptr1)
	//Get value through pointer
	fmt.Println("Pointer contains value is :", *ptr)

	/*

		//Two way
		//one way

		name := "Majharul"
		ptr2 := &name
		fmt.Println("name has value : ", name)
		fmt.Println("pointer contains : ", ptr2)
		//Get value through pointer
		fmt.Println("Pointer contains value is :", *ptr2)


		//second way
		var namePtr *string
		name := "Majharul"
		namePtr = &name

		fmt.Println("Pointer contains address:", namePtr)
		fmt.Println("Pointer contains value:", *namePtr)

	*/

	var pointer *int
	//Default pointer == nill
	//If we declare a int without defining int numbers then it stores automatically 0
	//Like this pointer store nill
	if pointer == nil {
		fmt.Println("Pointer is not assigned.")
	}

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value contains : ", value)

}
