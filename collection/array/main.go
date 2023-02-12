package main

import (
	"fmt"
)

func main() {
	// Declaring array
	var a [5]int

	// Initializing array
	a[0] = 10
	a[1] = 20

	fmt.Println("a:", a)

	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println("b:", b)
	fmt.Println(len(b)) // len is a built-in function

	// Array literal with ...
	c := [...]int{10, 20, 30, 40, 50}
	fmt.Println("c:", c)
	fmt.Println(len(c))

	// Initializing value at specific index using array literal.
	d := [3]int{0: 10, 2: 30}
	fmt.Println("d:", d)
	fmt.Println(len(d))

	// Iterate over array.
	for i := 0; i < len(d); i++ {
		fmt.Printf("%d\t", d[i])
	}
	fmt.Println()

	// Iterate over array using range.
	for i, v := range d {
		fmt.Printf("i:%d, v:%d\t", i, v)
	}
	fmt.Println()
}
