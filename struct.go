// This is a Go program demonstrating struct
// Author: Samuel Parente

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"Samuel", 30}
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
}
