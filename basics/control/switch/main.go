package main

import "fmt"

/*
switch condition{
	case expression1: Statement..
	case expression2: Statement..
	...
	default: Statement..
}
*/

func isVowel(ch string) bool {
	switch ch {
	case "a", "e", "i", "o", "u":
		return true
	default:
		return false
	}
}

func getType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Boolean")
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	{
		ch := "a"
		fmt.Printf("Is %s a vowel? %v\n", ch, isVowel(ch))
	}

	getType(true)
}
