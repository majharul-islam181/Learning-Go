package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println("Current time : ", currentTime)
	fmt.Printf("Type of current time : %T\n", currentTime) // Type of current time : time.Time

	// 2006-01-02 15:04:05
	formatted := currentTime.Format("02-01-2006, Monday, 15:04:05")
	fmt.Println("Formatted time : ", formatted) // Formatted time :  25-03-2025, Tuesday, 10:22:42

	// Time formating
	formatted1 := currentTime.Format("02/01/2006, Monday, 03:04 PM")
	fmt.Println("Formatted time : ", formatted1) // Formatted time :  25/03/2025, Tuesday, 10:22 AM

	// String to Time convertion
	layout_str := "2006-01-02"
	dateStr := "2023-11-01"
	formattedTime, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formated time : ", formattedTime)                               //  2023-11-01 00:00:00 +0000 UTC
	fmt.Println("Formated time with '/' : ", formattedTime.Format("02/01/2006")) //  01/11/2023
	fmt.Printf("Type of Formatted Time : %T\n", formattedTime)                   //  time.Time

	// Add 1 more day in currentTime
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("new date time : ", new_date)
	formatter_thisDate := new_date.Format("2006/02/01")
	fmt.Println("formated new date time; ", formatter_thisDate)
}
