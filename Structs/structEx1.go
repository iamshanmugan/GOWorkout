package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	shaan := person{firstName: "Shaan", lastName: "Nagarajan"}
	fmt.Println(shaan)

	raju := person{firstName: "shanmugam", lastName: "Nagarajan"}

	fmt.Println(raju)
	fmt.Println(shaan)
}
