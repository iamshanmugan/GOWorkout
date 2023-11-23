package main
import "fmt"

func main(){
	var arr[5] int
	arr = [...]int {5,4,3,2,1}
	for _,n:=range arr{
		fmt.Printf("%d\n",n)
	}
}