package main

import (
	"encoding/json"
	"fmt"
)

/*
	In Go, the encoding/json package is used to encode and decode Json data.
*/

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	City    string `json:"city"`
	IsAdult bool   `json:"isAdult"`
	IsMan bool       // automatic IsMan in jsonConvert
}

func main() {
	fmt.Println("Learning JSON in GOLANG")

	person := Person{Name: "Majharul", Age: 25, City: "Dhaka", IsAdult: true}
	// fmt.Println("Person data : ", person)

	// Convert struct to JSON
	// Convert person into JSON Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}
	fmt.Println("JSON data:", string(jsonData))

	// Decoding (UnMarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)

	if err != nil {
		fmt.Println("Error unmarshalling ", err)
		return
	}
	fmt.Println("person data is : ", personData)

}
