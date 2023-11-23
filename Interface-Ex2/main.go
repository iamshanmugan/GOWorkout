package main

import "fmt"

func main() {
	x := 235
	print(x)
	print(Employee{Name: "Shaan"})
	// if pointer using
	print(&Employee{Name: "Shaan"})
	print(Manager{title: "Technical Architect"})
}

type Employee struct {
	Name string
}

type Manager struct {
	title string
}

func (m Manager) Title() string {
	return m.title
}

type HasTitle interface {
	Title() string
}

func print(data interface{}) {
	switch val := data.(type) {
	case string:
		fmt.Printf("%s is a string \n", val)

	case int:
		fmt.Printf("%d is an integer \n", val)

	case Employee:
		fmt.Printf("%s is an Employee\n", val.Name)

	case *Employee:
		fmt.Printf("%s is an Employee from address\n", val.Name)

	case HasTitle:
		fmt.Printf("%s title is \n", val.Title())
	default:
		fmt.Printf("Unsupportive type....\n")
	}
}
