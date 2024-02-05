//Structs

package main

import "fmt"

type Client struct{
	Name string
	Age int
	Active bool
}

func main(){
	John := Client{
		Name: "John",
		Age: 33,
		Active: true,
	}
	fmt.Printf(" Name %s, age %d , active %t \n", John.Name, John.Age, John.Active)
	John.Active = false

	fmt.Println(John.Active)
}