// This is a Go program demonstrating switch statement
// Author: Samuel Parente

package main

import "fmt"

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("It's Monday!")
	case "Tuesday":
		fmt.Println("It's Tuesday!")
	default:
		fmt.Println("It's another day!")
	}
}
