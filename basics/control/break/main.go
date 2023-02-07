package main

import "fmt"

func main() {
	fmt.Println("====== Break =======")
	{
		for i := 0; i < 5; i++ {
			for j := 0; j < 10; j++ {
				if j == 5 {
					fmt.Printf("j is %d,\t", j)
					fmt.Println("Breaking inner loop")

					break
				}
			}
			fmt.Println("i:", i)
		}
	}

	fmt.Println("====== Break with lable =======")
	{
	OuterLoop:
		for i := 0; i < 5; i++ {
			for j := 0; j < 10; j++ {
				if j == 5 {
					fmt.Printf("j is %d,\t", j)
					fmt.Println("Breaking out of loop")
					break OuterLoop
				}
			}
		}
	}
}
