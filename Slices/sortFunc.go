package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"banana", "apple", "kiwi"}

	lenSort := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}
	slices.SortFunc(fruits, lenSort)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Shaan", age: 43},
		Person{name: "Kani", age: 10},
		Person{name: "Sai prasanna", age: 5},
		Person{name: "Vino", age: 37},
		Person{name: "Nagaraj", age: 67},
		Person{name: "VAlli", age: 65},
	}

	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(b.age, a.age)
		})
	fmt.Println(people)
}
