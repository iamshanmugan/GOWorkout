package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	shaan := person{firstName: "shaan", lastName: "Nagarajan"}
	fmt.Println(shaan)

	var raju person
	raju.firstName = "shanmugam"
	raju.lastName = "Nagarajan"
	fmt.Println(raju)

	sachu := person{
		firstName: "SaiPrasanna",
		lastName:  "Shanmugam",
		contact: contactInfo{
			email:   "sai@gmail.com",
			zipCode: 609999,
		},
	}
	//sachuPointer := &sachu
	sachu.updateName("Saiii")
	sachu.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName

}
