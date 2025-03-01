package main

import "fmt"

func main() {
	fmt.Println("Learning Array")
	/*
		//Array Initialization
		var name [3]string
		// 0 1 2 3
		name[0] = "prince"
		name[2] = "majharul"
		fmt.Println("Names of Person is : ", name)

		//Array Initialization
		var numbers = [5]int{1, 2, 3, 4, 5}
		fmt.Println("Number is : ", numbers)

		//Array length
		fmt.Println("Length of Numbers array is : ", len(numbers))

		fmt.Println("Value of name at 2 index : ", numbers[2])
	*/

	var price [5]int
	fmt.Println(price) // output: [0 0 0 0 0]

	var coming [5]bool
	fmt.Println(coming) // output: [false false false false false]

	var name [5]string
	fmt.Println(name) // output: [    ]
	//if you want to see the actual content of the array
	fmt.Printf("%q\n", name) // output: ["" "" "" "" ""] //formated way te dekhar jonno %q

	var office [5]string
	office[2] = "mombai"
	office[1] = "bonani"
	fmt.Printf("Format way of office : %q\n ", office) // output: ["" "bonani" "mombai" "" ""]

}
