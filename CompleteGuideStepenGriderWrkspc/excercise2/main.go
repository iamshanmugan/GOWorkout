package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args[1])

	fileDtls, err := os.Open(os.Args[1])
	if err != nil {
		println("Errror:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, fileDtls)

}
