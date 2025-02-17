package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("This is first time learning go")
	fmt.Println("this is second time learning go")
	//import package
	myutil.PrintMessage("jakanaka ")
	//variable
	var age int = 12
	var name string = "majharul"
	var isComing bool = false
	myutil.PrintMessage(name)
	//myutil.PrintMessage(age) //PrintMeassge needs string
	fmt.Println("your age is : ", age)
	fmt.Println("Am I coming today? : ", isComing)

	//another method to assign variable
	// here myname is declare and initialized using :=
	myname := "md. majharul islam"
	fmt.Println(myname)

	//Public variable (exported)
	var PublicVariable int = 11
	//PublicVariable is visible and can be assesed from others packages

	//private variable (unexported)
	var privateVariable int = 10
	//privateVariable is only visible withnin the same package.

	fmt.Println(privateVariable, PublicVariable)

	privateFunction()

}

// Public function (exported)
func PublicFunction() {
	fmt.Println("This is Public function; can be access from others packages")
}

// private function (unexported)
func privateFunction() {
	fmt.Println("This is private function; can not be access from others packages")
}
