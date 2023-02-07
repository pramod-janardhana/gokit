package main

import (
	"fmt"
)

func RecoverFromPanic() {
	defer func() {
		fmt.Println("Deferred call - 2")
		if e := recover(); e != nil {
			// e is the value passed to panic()
			fmt.Println("Recover with: ", e)
		}
	}()

	panic("Just panicking for the sake of example")
	fmt.Println("This will never be called")
}

func main() {
	fmt.Println("Starting to panic")
	RecoverFromPanic()
	fmt.Println("Program regains control after the panic recovery")
}
