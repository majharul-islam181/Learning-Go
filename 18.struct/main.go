package main

import "fmt"

/*
	In Go, a struct(Structure) is a composite data type that groups together variables
	(fields or members) under a single name.

	Each Filed in a struct can have a different data type, and structs are used to create
	more complex data structures.

*/

type Person struct {
	Firstname string
	Lastname  string
	Age       int
}

type contact struct {
	Email string
	Phone string
}

type address struct {
	HouseNumber int
	Area        string
	State       string
}

type Employee struct {
	Person_details Person
	Person_Contact contact
	Person_address address
}

func main() {

	var prince Person
	fmt.Println("Prince Person: ", prince)
	// Prince Person:  {  0}

	prince.Firstname = "Prince haidar"
	prince.Lastname = "Ali"
	prince.Age = 22
	fmt.Println("Prince all Details : ", prince)
	// Prince all Details :  {Prince haidar Ali 22}

	// 2nd Method to Initialize
	person1 := Person{
		Firstname: "Akash",
		Lastname:  "Islam",
		Age:       23,
	}
	fmt.Println("Details of person1 : ", person1)
	// Details of person1 :  {Akash Islam 23}

	// 3nd Method to Initialize
	// with *New* keyword
	var person2 = new(Person)
	person2.Firstname = "Simran"
	person2.Lastname = "Islam"
	person2.Age = 11

	fmt.Println("Details of Simran is : ", person2)
	// Details of Simran is :  &{Simran Islam 11}
	fmt.Println("Age of Simran is : ", person2.Age)
	// Age of Simran is :  11

	//Employee 1
	employee1 := Employee{
		Person_details: Person{
			Firstname: "Majharul",
			Lastname:  "Islam",
			Age:       11,
		},
		Person_Contact: contact{
			Email: "majharul.flutter@gmail.com",
			Phone: "01522211",
		},
		Person_address: address{
			HouseNumber: 121,
			Area:        "DHaka",
			State:       "Basundhara",
		},
	}

	fmt.Println("Details of employee1 : ", employee1)
	//Details of employee1 :  {{Majharul Islam 11} {majharul.flutter@gmail.com 01522211} {121 DHaka Basundhara}}
	fmt.Println("Name of employee1 : ", employee1.Person_details.Firstname)
	// Name of employee1 :  Majharul

}
