// This is a Go program demonstrating functions
// Author: Samuel Parente

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	result := add(10, 20)
	fmt.Println("Result:", result)
}
