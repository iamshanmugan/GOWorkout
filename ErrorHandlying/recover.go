package main

import "fmt"

func maypanic() {
	panic("a problem")
}
func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd Error..", r)
		}
	}()

	maypanic()
	fmt.Println("This wont call...After panic")

}
