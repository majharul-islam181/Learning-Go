package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, What's your name?")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.",name)

	reader := bufio.NewReader(os.Stdin)
	//creates a new buffered that reads from the standard input(os.Stdin)
	name, _ := reader.ReadString('\n')
	//reads a line from the input until it encounters a newline character('/n').
	//This allows the program to read the entire line of input, including spaces.
	fmt.Println("my name is : ", name)
	/*
		A buffered reader is a type of reader that reads data from an underlying souce,
		such as a file or standart input(keyboard), and store that data in a buffer.
		----------------------------------------------------------------------------
		The buffer is a temporary storage area in memory.
		Buffered readers are commonly used to improve the efficiency of input operations.

	*/

}
