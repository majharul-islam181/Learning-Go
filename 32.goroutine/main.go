package main

import (
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("Say hello is called")
	time.Sleep(10000 * time.Millisecond)
}

func sayHi(){
	fmt.Println("Say hi is called")
	
}
func main(){

	fmt.Println("Leaning go routine")
	go sayHello()
	sayHi()
	
time.Sleep(1000 * time.Millisecond)


}