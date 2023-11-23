package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("not work with this value...")
	}
	return arg + 3, nil

}

type argError struct {
	arg  int
	prob string
}

// This is the custom type error implementation.
func (e *argError) Error() string {
	return fmt.Sprintf("%d, %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "not workwith this valu... in f2"}
	}
	return arg + 3, nil
}

func main() {
	//This is using the errors type
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed ", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}

	//This is using the custom errors type
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed ", e)
		} else {
			fmt.Println("f2 worked..", r)
		}
	}

	//handling programatically the custom error
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
