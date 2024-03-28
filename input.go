// This is a Go program demonstrating user inputs
// Author: Samuel Parente

package main

import "fmt"

func main() {
	// Taking input for a single variable
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello,", name)

	// Taking input for multiple variables
	var age int
	var city string
	fmt.Print("Enter your age and city separated by space: ")
	fmt.Scanln(&age, &city)
	fmt.Println("Age:", age, "City:", city)
}
