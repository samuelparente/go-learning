// This is a Go program demonstrating file handling
// Author: Samuel Parente

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, file handling!")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File created and written successfully.")
}
