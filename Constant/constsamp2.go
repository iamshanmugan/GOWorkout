package main

import (
	"fmt"
)

// Day of week type
type DayofWeek int

const (
	//Monday DayofWeek = iota
	//Monday DayofWeek =1 + iota
	Monday DayofWeek = 1 << iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println(Sunday)
}
