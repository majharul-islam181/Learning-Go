package main

import "fmt"

func main() {
	/*
		day := 6

		switch day {
		case 1:
			fmt.Println("Staurday")
		case 2:
			fmt.Println("Monday")
		case 3:
			fmt.Println("Thursday")
		default:
			fmt.Println("Unkown Day")
		}

	*/

	day := 7

	switch day {
	case 1, 2, 3, 4, 5, 6, 7:
		fmt.Println("1st week")
	case 8:
		fmt.Println("Another week")
	default:
		fmt.Println("Unkown weeks ")
	}
}
