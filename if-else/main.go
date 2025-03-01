package main

import "fmt"

func main() {
	x := 2
	if x > 9 {
		fmt.Println("x is greater then 9")

	} else {
		fmt.Println("x is smaller than 5")
	}

	z := 10

	if z > 10 {
		fmt.Println("z is greater then 10")
	} else if z < 10 {
		fmt.Println("z is lower the 10")
	} else {
		fmt.Println("Nothing will print here")
	}

	y := 10

	if y > 5 && y != 0 {
		fmt.Println("This is working fine.......")
	}
}
