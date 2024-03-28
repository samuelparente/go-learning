// This is a Go program demonstrating JSON handling
// Author: Samuel Parente

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "Samuel", Age: 30}

	// Encoding to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData))

	// Decoding from JSON
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Decoded Person:", decodedPerson)
}
