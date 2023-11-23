package main

import "fmt"

func main() {
	color := map[string]string{
		"red":   "#0001",
		"green": "#0002",
	}

	//color := make(map[int]string)
	//fmt.Println(color)

	//color[10] = "#FFFF"
	//fmt.Println(color)

	//delete(color, 10)
	//fmt.Println(color)
	printMap(color)

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}

}
