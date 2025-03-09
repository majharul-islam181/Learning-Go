package main

import (
	"fmt"
)

func main() {
	println("Learning for loop")

	for i := 0; i < 1; i++ {
		fmt.Println("Number is : ", i)
	}

	// There is no do-while loop in Go, but you can achieve similar behavior using a for loop with a conditional
	// check a the end.

	counter := 0
	for {
		fmt.Println("Infinite Loop : ", counter)
		counter++
		if counter == 1 {
			break
		}

	}

	// numbers := []int{1, 2, 3, 4, 5}
	// for index, value := range numbers {
	// 	fmt.Printf("Index of Numbers is %d and value is %d\n", index, value)
	// }

	//Range in String
	data := "Hello world"

	for index, char := range data {
		fmt.Printf("Index of data is %d and value is %c\n", index, char)
	}

}
