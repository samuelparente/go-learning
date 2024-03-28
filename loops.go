// This is a Go program demonstrating loops
// Author: Samuel Parente

package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i:", i)
	}

	// while loop equivalent
	j := 0
	for j < 5 {
		fmt.Println("Value of j:", j)
		j++
	}
}
