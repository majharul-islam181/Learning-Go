package main

import (
	"fmt"
	"strings"
)

func main(){
	data := "apple, orange, banana"
	parts := strings.Split(data,",")
	fmt.Println("datat is : ", parts)


	str := "one two three four two two five five"
	cout := strings.Count(str, "two")
	fmt.Println("Count is : ", cout)


	str = "    Heloo World!                           ";
	trim := strings.TrimSpace(str)
	fmt.Println("Trim space : ", trim)


	str1 := "Prince"
	str2 := "Hudai"
	result := strings.Join([]string{str1, str2}, ",")
	fmt.Println("result is : ",result)
}