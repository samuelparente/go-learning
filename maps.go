// This is a Go program demonstrating maps
// Author: Samuel Parente

package main

import "fmt"

func main() {
	ages := map[string]int{
		"Samuel": 30,
		"Alice":  25,
		"Bob":    35,
	}

	fmt.Println("Age of Samuel:", ages["Samuel"])
}
