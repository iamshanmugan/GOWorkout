package main
import (
	"fmt"
	"io"
	"os"
)

func main(){
readFile(os.Args[1])
}

func readFile(name string){

	var bytes [10] byte
	f,err := os.Open(name)
	defer f.Close()
	if err!=nil{
		panic(err)
	}
	
	for{
		_,err :=f.Read(bytes[:])
		if err==io.EOF{
			break
		}
		fmt.Printf("%s",bytes)
	}
}