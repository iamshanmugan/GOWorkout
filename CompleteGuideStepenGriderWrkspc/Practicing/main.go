package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

// func main() {
// 	fmt.Println("Hi Raju - This is Testing practice by me...")
// 	//START - to print even or odd of the slice values

// 	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	// for _, number := range numbers {
// 	// 	if number%2 == 0 {
// 	// 		fmt.Println("EVEN ---", number)
// 	// 	} else if number%2 != 0 {
// 	// 		fmt.Println("ODD ---", number)
// 	// 	}

// 	// }

// 	//END - to print even or odd of the slice values
// 	//START - to simple STRUCT to use

// 	type contactInfo struct {
// 		email    string
// 		contatno int
// 	}

// 	//raju := person{"shanmugam", "nagarajan"}
// 	shan := person{
// 		firstName: "shanmugam",
// 		lastName:  "nagarajan",
// 	}
// 	rajuPointer := &shan

// 	rajuPointer.updateFirstname("Raju")

// 	chking := "shaan"
// 	fmt.Println(*&chking)
// 	shan.print()

// 	name := "bill"

// 	namePointer := &name

// 	fmt.Println(&namePointer)
// 	printPointer(namePointer)
// }

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
	//END - to simple STRUCT to use
}

func (p person) print() {
	fmt.Println(p)
}

func (pointerofPErson *person) updateFirstname(newfirstname string) {

	(*pointerofPErson).firstName = newfirstname
}
