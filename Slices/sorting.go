package main

import (
	"fmt"
	"slices"
)

func main() {
	str := []string{"d", "c", "a"}
	slices.Sort(str)
	fmt.Println("sorted str: ", str)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("sorted ints: ", ints)

	s := slices.IsSorted(ints)
	fmt.Println(s)

}
