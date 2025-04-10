package main

import (
	"fmt"
	"net/url"
)

func main(){

	fmt.Println("Learning url")

	myUrl := "https://facebook.com/newPath?key=iammajharul"

	fmt.Printf("myURL type is : %T\n",myUrl)

	parseUrl, error := url.Parse(myUrl) // String to parse into URL 
	if error != nil{
		fmt.Println("Cannot parse URL")
	}

	// fmt.Printf("Type of URL, %T\n", parseUrl)

	fmt.Println("Schema: ",parseUrl.Scheme)
	fmt.Println("Host : ", parseUrl.Host)
	fmt.Println("Path : ",parseUrl.Path)
	fmt.Println("RawQuery : ",parseUrl.RawQuery)

	// Modifing URL components
	parseUrl.Path = "/newpath"
	parseUrl.RawQuery = "username=majharul"

	// Constructing a URL string from a URL object
	newURl := parseUrl.String()
	fmt.Println("new url : ",newURl)

}