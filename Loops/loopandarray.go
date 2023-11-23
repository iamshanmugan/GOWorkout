package main
import (
	"fmt"
	"strconv"
)

func main(){

	// this is used as for loop
	var temp,sum int
	// for i:=0; i<5; i++{
	// 	fmt.Scanf("%d\n",&temp)
	// 	sum = temp + sum
	// }
	// fmt.Printf("sum is %d ",sum)

	// this is used as while  loop but syntax is using for loop
	// i:=0
	// for i<5 {
	// 	fmt.Scanf("%d\n",&temp)
	// 	sum = temp + sum
	// 	i++
	// }
	// fmt.Printf("sum is %d ",sum)


	// this is used as do while  loop but syntax is using for loop
	var str string
	var err error
	for{
		fmt.Scanf("%s\n", &str)
		temp,err =  strconv.Atoi(str)
		if err!=nil {
			break
		}
		sum = temp + sum
	}
	fmt.Printf("sum is %d ",sum)

}