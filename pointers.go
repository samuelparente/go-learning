// This is a Go program demonstrating pointers
// Author: Samuel Parente

package main

import "fmt"

func main() {
	x := 10
	ptr := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Value at pointer:", *ptr)
}
