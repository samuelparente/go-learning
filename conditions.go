// This is a Go program demonstrating conditions
// Author: Samuel Parente

package main

import "fmt"

func main() {
	x := 10
	y := 20

	if x > y {
		fmt.Println("x is greater than y")
	} else if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is equal to y")
	}
}
