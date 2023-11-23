package main
import(
	"fmt"
	"unicode"
	"strconv"
) 

func main(){
	var subval byte
	fmt.Println("Do you want to subscribe (y/n)? ")
	fmt.Scanf("%c",&subval)

	subval = byte(unicode.ToLower(rune(subval)))
	switch subval{
	case 'y', 'Y':
		fmt.Println("Thank you for subscrition...")
		fallthrough
	
	case 'n', 'N':
		fmt.Println("Panatha unaku thaan loss....")
	
	default:
	fmt.Println("Invalid Input...")
	}

	fmt.Printf("%s",fizzBuzz(15))
}


func fizzBuzz(n int) string{

	switch {
	case n%3 ==0 && n%5==0:
			return "Fizz"
			
	
	case n%3==0:
			return "Buzz"

	case n%5==0:
		return "FizzBuzz"
	default:
		return strconv.Itoa(n)
	}	

}