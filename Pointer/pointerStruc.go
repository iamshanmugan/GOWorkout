package main

import "fmt"

func main() {
	var a, b int
	//to declar as pointer see the c var
	var c *int
	a = 10
	b = a
	b = 5
	c = &b

	fmt.Printf("avalue --> %d \n   bvalue --> %d\n    cvalue is pointer ---> %d\n      c pointer address --->%d \n ", a, b, *c, c)

	// this below function is using the swap meth below to swap the 2 value using pointer.

	d, e := 10, 15
	swap(&d, &e)
	fmt.Printf("%d  %d", d, e)

}

func swap(d, e *int) {
	*d, *e = *e, *d
}
