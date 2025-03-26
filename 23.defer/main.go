package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Starting of the programm")
	sum := add(5, 6)
	defer fmt.Println("The sum is : ", sum)
	defer fmt.Println("Middle of the programm")
	defer fmt.Println("End of the programm")

	fmt.Println("Very end of the programm")
}

/*
	Starting of the programm
	Very end of the programm
	End of the programm
	Middle of the programm
	The sum is :  11
*/
