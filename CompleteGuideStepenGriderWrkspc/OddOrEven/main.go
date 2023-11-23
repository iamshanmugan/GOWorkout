package main

import "fmt"

func main() {
	fmt.Println("testing")

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, num := range numbers {
		fmt.Println("index is %v, and number in the array is %v", i, num)

		if num%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
