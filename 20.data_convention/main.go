package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 300
	fmt.Println("num is ", num)
	fmt.Printf("Type is num is : %T\n ", num)

	//Int to Float
	var data float64 = float64(num)
	data = data + 1.2
	fmt.Println("data is: ", data)
	fmt.Printf("Type of data is (Int to Float) : %T\n", data) //Type of data is (Int to Float) : float64

	//Int to string
	str := strconv.Itoa(num)
	fmt.Println("Sting is : ", str)
	fmt.Printf("Type of str is (Int to String) : %T\n", str) //Type of str is (Int to String) : string

	//String to Number
	number_string := "121212"
	num_int,_ := strconv.Atoi(number_string);
	fmt.Println("Value of number_string : ",num_int)
	fmt.Printf("Num_int Type is(String to Int): %T\n",num_int) //Num_int Type is(String to Int): int



	//String to Float Number
	num_string := "12.1212"
	num_float,_ := strconv.ParseFloat(num_string, 64);
	fmt.Println("Value of number_string : ",num_float)
	fmt.Printf("Num_float Type is(String to Int): %T\n",num_float) //Num_int Type is(String to Int): int
}
