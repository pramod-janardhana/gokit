package main

import "fmt"

func main() {
	i := 0
skip:
	for i < 10 {
		if i == 5 {
			fmt.Println("Skippinng value", i)
			i++
			goto skip // rebirect the control here
		}
		fmt.Println("Value of i is", i)
		i++
	}
}
