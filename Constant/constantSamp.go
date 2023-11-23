package main
import(
	"fmt"
	"reflect"
)
const(
	X=1
	Y=2
)

type Farenheit float64
type Celcius float64

// DefaultUnit is a type alias
type DefaultUnit =   Celcius
func main(){
	const PI =2.322212
	fmt.Printf("%f\n",PI)

	 var a float32
	 var b float64


	// var r Farenheit
	 var s Celcius
	 var t DefaultUnit

	// r=100
	 s=32

	
	 t=s
	 //fmt.Printf("%f\n",t)
	 fmt.Println(reflect.TypeOf(t))
	 a=PI
	 b=PI
	 fmt.Printf("%f\n",a)
	 fmt.Printf("%f\n",b)


	 
}