package main

import "fmt"

func main() {
	weekdays := [...]string{
		"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN",
	}
	workingDay := weekdays[1:6]
	workingDay[2] = "Wednesday"
	printslice(workingDay)
	printslice(weekdays[:])

	// adding the value in the arrays
	var list []int
	listB := make([]int, 5, 10)
	list = append(list, 1, 2)
	printsliceint(list)
	list = append(list, 3)
	printsliceint(list)
	printsliceint(listB)

	var copyListA []int

	copyListA = append(copyListA, 1, 2)
	copyListB := make([]int, len(copyListA))
	copy(copyListB, copyListA)
	copyListB[0] = 3
	printsliceint(copyListA)
	printsliceint(copyListB)
}

func printslice(slice []string) {
	fmt.Printf("%v\n", slice)
}

func printsliceint(slice []int) {
	fmt.Printf("%v, Len: %d, cap: %d \n", slice, len(slice), cap(slice))
}
