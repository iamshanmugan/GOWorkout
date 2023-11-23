package main
import(
	"fmt"
)
func main (){

	x , y := div(15,10)
	fmt.Printf("%d, %d\n", x, y)


	//using the function as value
	square:=func(a int) int{
		return a*a
	}
	fmt.Printf("square value of  ")
	fmt.Println(square(2))


	// pass the function as a parameter
	fmt.Println(apply(4,square))

	// using the function as return type
	fn := powerMux(3)
	u := fn(4)
	fmt.Println(u)

	//variodic func which accept multiple parameter without defining fixed.
	resultToPrint := sum(1,2)
	fmt.Printf("%d\n",resultToPrint)


	// defer function used to do some operation after the function call executed.
	doSomething()
}

func doSomething(){
	defer fmt.Printf("Done")

	defer func(){
		fmt.Printf("Donnnnnnn")
	}()
	fmt.Printf("Doing")
}

func sum(n ...int) int{
	result :=0
	for _, a:=range n{
		result += a
	}
	return result
}

func apply(a int, fn func(int)int)int{
	return fn(a)
}

func add(a int, b int) (c int){
	c=a+b
	return
}

func div(num int , denom int)(int , int){
	return num/denom, num % denom
}


func powerMux(n int) func(int)int{
	switch n{
	case 2:
		return func(a int)int {return a*a}
	case 3:
		return func(a int)int {return a*a*a}
	default :
		return func(a int)int {return 0}
	}
}

