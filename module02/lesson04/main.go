// Working with json.

package main

import (
	"encoding/json"
	"os"
)

type Account struct{
	Number int `json:"number"`
	Ammount int `json:"ammount"`
	Client string `json:"ammount"`
}

func main(){
 account := Account{Client: "John", Number: 1, Ammount: 100}
 result, err := json.Marshal(account)
 if err!= nil {
	panic(err)
 }
 println(string(result))

 err = json.NewEncoder(os.Stdout).Encode(account)
 if err!= nil {
	println(err)
 }

  purejson :=[]byte(`{"Number":2,"Ammount":200,"Client":"Petter"}`) 
  var acount_x Account
  err=json.Unmarshal(purejson,&acount_x)
  if err != nil {
	println(err)
  }
  println(acount_x.Ammount)
}