// This is a Go program demonstrating goroutines and concurrency
// Author: Samuel Parente

package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	go printNumbers()
	time.Sleep(5 * time.Second)
}
