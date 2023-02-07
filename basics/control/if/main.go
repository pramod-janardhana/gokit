package main

import "fmt"

/*
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ]

if  condition-1 {
    // code to be executed if condition-1 is true
} else if condition-2 {
    // code to be executed if condition-2 is true
} else {
    // code to be executed if both condition1 and condition2 are false
}
*/

func isEven(input int) {
	if input%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}

func main() {
	isEven(25)
}
