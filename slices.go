// This is a Go program demonstrating slices
// Author: Samuel Parente

package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	slice := numbers[1:4]
	fmt.Println("Slice:", slice)
}
