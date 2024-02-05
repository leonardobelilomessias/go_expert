// Pointers and structs

package main

import "fmt"

type Client struct{
	Name string
}

func newClient()*Client{
	return &Client{Name: "new client"}
}

func(c *Client)walk(){ //with * we pass pointer of struct and change value in adderess
	c.Name = "John Dove"
	fmt.Printf("The client %v andou\n", c.Name)
}
func main(){
	john := Client{
		Name: "John",
	}
	john.walk()
	fmt.Printf("The value the struc with name %v\n",john.Name)
	myclient := newClient()
	fmt.Println(&myclient)
}