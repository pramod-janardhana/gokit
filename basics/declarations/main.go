package main

import "fmt"

/*
https://www.w3schools.com/go/go_data_types.php
Basic types
1. Number(Integer, Float)
2. String
3. Boolean
*/

// iota can used to declare enums
const (
	// TRACE logs everything
	TRACE int = iota // 0
	// INFO logs Info, Warnings and Errors
	INFO // 1
	// WARNING logs Warning and Errors
	WARNING // 2
	// ERROR just logs Errors
	ERROR // 3
)

// const declares a constant value.
const c string = "I'm a constant"

func main() {
	// use "var" keyword to declare 1 or more variables.
	var firstname, lastname string
	var age int
	fmt.Printf("firstname: %s, lastname: %s, age: %d\n", firstname, lastname, age)

	// 1. declare and initialize variables with type.
	var a1 int = 1
	fmt.Println(a1)

	// 2. declare and initialize variables without type.
	var a2 = 2
	fmt.Println(a2)

	// 3. declare and initialize variables without "var" keyword.
	// The := syntax is shorthand for declaring and initializing a variable.
	a3 := 3
	fmt.Println(a1, a2, a3)
}
