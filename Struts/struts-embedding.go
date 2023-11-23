package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num %v", b.num)
}

type container struct {
	base
	str string
}

func main() {

	b := base{
		num: 1,
	}
	fmt.Println(b)
	fmt.Println(b.describe())

	co := container{
		base: base{
			num: 1,
		},
		str: "Shaan",
	}
	fmt.Println(co.base.num)
	fmt.Println(co.num)
	fmt.Println(co.str)
	fmt.Println(co.describe())

	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer ", d.describe())

}
