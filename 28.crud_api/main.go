package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json: "UserId"`
	Id        int    `json: "id"`
	Title     string `json: "title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, error := http.Get("https://jsonplaceholder.typicode.com/tod")
	if error != nil {
		fmt.Println("Error getting , ", error)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Eror iin getting response : ", res.StatusCode)
		// fmt.Println("Eror iin getting response : ", res.Status)

		return
	}
	// Api data read korar jonno
	// data, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Println("Error reading : ", err)
	// 	return
	// }
	// fmt.Println("Data : ", string(data))

	var todoData Todo
	err := json.NewDecoder(res.Body).Decode(&todoData)
	if err != nil {
		fmt.Println("Error decoding : ", err)
		return

	}
	fmt.Println("TODO : ", todoData)

	fmt.Println("Title response : ", todoData.Title)
	fmt.Println("Completed response : ", todoData.Completed)

}

func performPostRequest() {

	todo := Todo{
		UserId:    22,
		Id:        3,
		Title:     "Majharul",
		Completed: false,
	}

	// Convert the TODO struct into JSON
	jsonData, error := json.Marshal(todo)
	if error != nil {
		fmt.Println("Error marshalling : ", error)
		return
	}
	// Convert JSON data to String
	jsonString := string(jsonData)

	// Convert String to Reader
	jsonReader := strings.NewReader(jsonString)

	myURl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myURl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error in POST METHOD ", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("RESPONSE DATA : ", res)

	// (res *http.Response) Response Reading
	response, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response of BODY : ", response)

}

func performUpdateRequest() {
	todo := Todo{
		UserId:    22,
		Id:        3,
		Title:     "Majharul",
		Completed: false,
	}

	// Data to JSON with Marshal
	res, error := json.Marshal(todo)
	if error != nil {

		fmt.Println("JSON error : ", error)
		return
	}
	fmt.Println("RESPONSE : ", string(res))

	// convert json data(byte) to string
	jsonString := string(res)

	// convert STRING to READER
	jsonReader := strings.NewReader(jsonString)

	myURl := "https://jsonplaceholder.typicode.com/todos"

	// Create PUT REQUEST
	req, er := http.NewRequest(http.MethodPut, myURl, jsonReader)

	if er != nil {
		fmt.Println("Error createing put requet: ", er)
		return
	}
	req.Header.Set("Content-type", "application/json")

	// Send Request
	client := http.Client{}
	ress, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}
	defer ress.Body.Close()

	data, _ := ioutil.ReadAll(req.Body)
	fmt.Println("Response : ", string(data))
	fmt.Println("Response status : ", ress.Status)

}

func performDeleteRequest() {

	const myURl = "https://jsonplaceholder.typicode.com/todos"
	// Create DELETE REQUEST

	req, error := http.NewRequest(http.MethodDelete, myURl, nil)
	if error != nil{
		fmt.Println("Error in DELETE REQUEST", error)
		return
	}

	//Sent the request
	client := http.Client{}
	res, er := client.Do(req)
	if er != nil{
		fmt.Println("Error in delete request", er)
		return
	}
	defer res.Body.Close()


	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response data : ",data)
	fmt.Println("Response status : ", res.Status)


}

func main() {
	fmt.Println("Learning CRUD API")

	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()

}
