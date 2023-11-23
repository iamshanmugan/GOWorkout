package main
import (
	"fmt"
	"encoding/json"
)
func main(){
	wheelsMap := map[string]int{
		"car":4,
		"auto":3,
		"bike":2,
	}
	//adding in the map.
	wheelsMap["suv"]=4

	//delete in the map
	delete(wheelsMap,"bike")



	fmt.Println(wheelsMap)


	//using for loop:
	for vehicle, wheel := range wheelsMap{
		fmt.Printf("%s, %d\n",vehicle,wheel)

	}


	// check if the given key is in map or not
	var input string
	fmt.Scanf("%s",&input)
	wheel,ok:=wheelsMap[input]
	if !ok{
		fmt.Printf("Wheel map not found")
		return
	}
	fmt.Printf("No of wheel is: %d",wheel)

	// converting the map obj to json
	data, _ := json.Marshal(wheelsMap)
	fmt.Printf("%s",data)
}