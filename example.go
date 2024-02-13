package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func main(){
	number := 10
	println(&number)
	number3 := number
	number2 := &number
	println(number2, number)
	*number2 = 6
	number3 =15
	println(*number2, number, number3)
	example(&number3)
	person := example2("John", 35)
	println("person example" , example2("John", 35))
	person.Age = 15
	person2 := Person{

	}
	fmt.Println("person2", person2)
	fmt.Println("person", person)
}
func example(n *int) *int{
	println("example",n)
	return n

}
func example2(name string , age int ) *Person{
	return &Person{
		Name: name,
		Age: age,
	}

}