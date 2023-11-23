package main
import(
	"fmt"
)
type Celcius float32
type Farenheit float32


func main(){
var c Celcius
c=36.9
f:=c.ToFarenheit()
fmt.Println(f)
}


func (c Celcius) ToFarenheit()Farenheit{
	return Farenheit(9.0/5.0 *c + 32)
}

func (f Farenheit) String()string{
	return fmt.Sprintf("%.2f%cF",f,0x00B0)
}