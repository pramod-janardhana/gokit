package main

import "fmt"

func twoSum(a, b int) int {
	return a + b
}

func nSum(arr ...int) (sum int) {
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return
}

func main() {
	fmt.Println(twoSum(1, 2))
	fmt.Println(nSum(1, 2, 3, 4))
	fmt.Println(nSum(1, 2, 3, 4, 5))
}
