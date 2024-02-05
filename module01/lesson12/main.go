//Structs compositions

package main

import "fmt"

type Adderess struct{
	City string
	Street string
	Number int
}

type Client struct{
	Name string
	Age int
	Active bool
	Adderess 
}

func main(){
	John := Client{
		Name: "John",
		Age: 33,
		Active: true,
	}
	fmt.Printf(" Name %s, age %d , active %t \n", John.Name, John.Age, John.Active)
	John.Active = false
	John.Adderess.City="Rio de Janeiro"
	fmt.Println(John.Active)
	fmt.Println(John)
}