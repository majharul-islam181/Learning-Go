package main

import "fmt"

func main(){

	// name => grade
	studentsGrade := make(map[string]int)
	studentsGrade["Majharul"] = 90
	studentsGrade["Chamok"]= 30
	studentsGrade["Islam"] = 200

	fmt.Println("Marks of majharul : ",studentsGrade["Majharul"])

	studentsGrade["Majharul"]= 100
	fmt.Println("New marks of Majharul : ", studentsGrade["Majharul"])


	//for deleting value of a map
	delete(studentsGrade, "Islam")
	fmt.Println("Marks of Islam person is : ", studentsGrade["Islam"]) //Marks of Islam person is :  0


	//Checking if a key exists or not
	values, exist := studentsGrade["Islam"]
	fmt.Println("Values are : ", values)
	fmt.Println("exist ?  ", exist) // return true/false


	for index,value := range studentsGrade{
		fmt.Printf("key is %s and value is %d\n", index, value)
	}



	// Another way to create map

	person_data := map[string]int{
		"majharul" : 12,
		"islam": 11,
	}

	for index,value := range person_data{
		fmt.Printf("key is %s and value is %d\n", index, value)
	}





}