package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("punch peach"))
	fmt.Println("r idx: ", r.FindStringIndex("punch peach"))
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	fmt.Println(r.Match([]byte("peach")))
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp: ", r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruits>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
