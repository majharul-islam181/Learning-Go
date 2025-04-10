package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")

	res, error := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if error != nil {
		fmt.Println("Error getting Get request")
		return
	}
	defer res.Body.Close()

	fmt.Printf("Type of response : %T\n", res)
	//    fmt.Println("response : ",res)

	data, error := ioutil.ReadAll(res.Body)
	if error != nil {
		fmt.Println("Error getting Get request")
		return

	} 

	fmt.Println("data : ", string(data))


}
