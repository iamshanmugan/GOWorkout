package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//This is the get the input from commmand line and print it
	//s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = ""

	// }
	// fmt.Println(s)

	//This is second program which is to get from command line and check for duplicates

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Println("test")
		if n > 1 {
			fmt.Println("n>1")
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
