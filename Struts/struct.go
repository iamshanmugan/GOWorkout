package main
import(
	"fmt"
	"time"
	"encoding/json"

)
	

type Customer struct{
	ID int `json:"id"`
	name string `json:"name"`
	DOB time.Time  `json:"dob"`
}

type Transaction struct{
	ID int `json:"id"`
	Time time.Time `json:"time"`
	Amount int `json:"amount"`
	Credit bool `json:"credit"`
}

type Account struct{
	ActNo int `json:"actno"`
	Holder *Customer `json:"holder"`
	Transaction []Transaction `json:"transaction"`
}

func main(){
	var c Customer
	c.ID=1
	c.name="Shaan"
	c.DOB=time.Date(1980, time.April, 22 ,0,0,0,0,time.Local)


	// only some value to set
	c= Customer{
		ID:2,
		name:"Raju",
	}
	fmt.Printf("%v", c)

	var acc Account
	acc=Account{
		ActNo:1,
		Holder:&Customer{
			ID:1,
			name:"Shaan",
		},
		Transaction:[]Transaction{
			{1,time.Now(),100,true},
			{2,time.Now(),100,false},
			
		},
		
	}
	data, err := json.Marshal(acc)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("json format %s\n",data)
	fmt.Printf("AcctDetails   : %v",acc)
}