package main

import (
	"fmt"
)

func main() {
	// Declares a nil map
	var chapters map[int]string

	// Initialize map with make function
	chapters = make(map[int]string)

	// Add data as key/value pairs
	chapters[1] = "Beginning Go"
	chapters[2] = "Go Fundamentals"
	chapters[3] = "Structs and Interfaces"

	// Iterate over the elements of map using range
	for k, v := range chapters {
		fmt.Printf("Key: %d Value: %s\n", k, v)
	}

	// Declare and initialize map using map literal
	languages := map[string]string{
		"EL": "Greek",
		"AR": "Arabic",
		"EN": "English",
		"ES": "Spanish",
		"FR": "French",
	}

	// Delete an element
	delete(languages, "EL")
	// Lookout an element with key
	if lan, ok := languages["EL"]; ok {
		fmt.Println(lan)
	} else {
		fmt.Println("\nKey doesn't exists")
	}

	// Update an existing item of map
	languages["EN"] = "US English"
	for k, v := range languages {
		fmt.Printf("Key: %s Value: %s\n", k, v)
	}
}
