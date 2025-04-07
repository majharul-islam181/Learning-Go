package main

import (
	"fmt"
	"io"
	"os"
	// "strings"
)

/*
File handling in Go invole working with the os and io/ioutil package.
Here a basic guide to common file operations like creating, reading and
writing files.
*/
//  
/*
we can use the io package in GO for writing to a file. The io.WriteString function
is a convenient wasy to write a string to a fie.

````

file.Close() is important in many cases when you are done working with a file.
When you open a file using functions like os.Create,os.Open, or others, you are acquiring system
resources to intereact with that file.
*/
func main() {

	/*
		// Create A File
		file, err := os.Create("example.text")
		if err != nil {
			fmt.Println("Error while creating file:  ", err)
			return
		}
		defer file.Close()

		// Write content into File
		content := "hello world"
		_, error := io.WriteString(file, content+"\n")
		if error != nil {
			fmt.Println("Error while creating file:  ", error)
			return
		}

		fmt.Println("successfully created a file")
	*/

	//File Open

	file, error := os.Open("example.text")
	if error != nil {
		fmt.Println("Error while opening file", error)
		return
	}
	defer file.Close()
	fmt.Println("successfully opened a file")

	// Create a buffer to read file content

	buffer := make([]byte, 1024)

	// Read the file content into the buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error while reading file", err)
			return
		}

		// Process the read content
		fmt.Println(string(buffer[:n]))

	}

}
