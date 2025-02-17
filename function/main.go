package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printSomething() {
	fmt.Println("Print something.........")
}

func addition(a, b int) (result int) {
	result = a + b
	return
}

func multiply(a int, b int) int {
	return a * b
}

func personData(name string, age int, salary float64) (string, int, float64) {
	return name + " extra added ", age + 10, salary + 1000
}

//in a Go application, the main function is the entry point of the programm.
//this main function does not take any argument and does not any return type

func main() {
	fmt.Println("We are learning function in golang")
	printSomething()
	sum := addition(2, 3)
	fmt.Println("sum is :", sum)
	mul := multiply(3, 4)
	fmt.Println("Multiplication : ", mul)

	//input from user
	fmt.Println("Enter your name ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Remove any newline character

	fmt.Println("Enter your age: ")
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age) // Remove any newline character
	age1, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("invalid age input")
	}

	fmt.Println("Enter your salary: ")
	salary, _ := reader.ReadString('\n')
	salary = strings.TrimSpace(salary) // Remove any newline character
	salary1, err := strconv.ParseFloat(salary, 64)
	if err != nil {
		fmt.Println("Invalid salary input")
	}

	name2, age2, salary2 := personData(name, age1, salary1)
	fmt.Printf("name is: %s, age is : %d, salary is: %.1f\n", name2, age2, salary2)

}
