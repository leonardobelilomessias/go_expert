//Maps

package main

import "fmt"

func main(){
	salaries := map[string]int{"wesley":1000, "john":2000}
	fmt.Println(salaries["wesley"])
	salaries["Peter"] = 2500
	games := make(map[string]int)
	games["fifa"] = 1990
	computers := map[string]int{}
	computers["intel"]=2000   
	fmt.Println(salaries)

	for name, salary:=range salaries{
		fmt.Printf(" The salary of %s is %d\n", name, salary)
	}
	for _, salary:=range salaries{
		fmt.Printf(" The salary is %d\n" , salary)
	}
}