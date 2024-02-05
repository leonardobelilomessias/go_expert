//Interfaces

package main

import "fmt"

type Adderess struct{
	City string
	Street string
	Number int
}

type Person interface{
	Deactive()
}

type Bussiness struct{
	Name string
}

func(b Bussiness)Deactive(){
	fmt.Printf("The bussiness %s was deactived  \n",b.Name,  )
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

func Deactivation(person Person){
	person.Deactive()
}

func main(){
	John := Client{
		Name: "John",
		Age: 33,
		Active: true,
	}
	myBussiness := Bussiness{}
	Deactivation(myBussiness)
	Deactivation(John)
}