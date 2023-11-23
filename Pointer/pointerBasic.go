package main

import "fmt"

func main() {
	fmt.Println("Go pointer learning....")

	var name string
	name = "shaan"
	fmt.Println(name)

	ptr := &name
	fmt.Println(ptr)

	*ptr = "Raju"

	fmt.Println(name)
}
