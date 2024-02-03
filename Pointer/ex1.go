package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2 // changing the x value using the address of the x
	fmt.Println(x)

	// to show how the pointer points the address(result: true,false,false)
	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil)

}
