package main

import "fmt"

/*
for initialization; condition; post{
       // statements....
}
*/
func sunOfNNumbers(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)
}

func printArray(arr []int) {
	for idx, val := range arr {
		fmt.Println(idx, val)
	}
}

func main() {
	sunOfNNumbers(3)

	{
		arr := []int{1, 2, 3, 4}
		printArray(arr)
	}
}
