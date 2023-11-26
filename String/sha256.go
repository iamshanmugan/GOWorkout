package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	s := "Sha256 this stirng"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
