//Structs methods

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

func (c Client) Deactive(){
	c.Active = false
	fmt.Printf("The client %s was deactived %t \n",c.Name, c.Active )
}

func main(){
	John := Client{
		Name: "John",
		Age: 33,
		Active: true,
	}

	fmt.Println(John.Active)
	John.Deactive()
	fmt.Println(John.Active)
	fmt.Println(John)
}