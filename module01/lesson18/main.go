// Empyt interface

package main

import "fmt"

type x interface{

}

func main(){
	var x interface{} = 10
	var y interface{} = "hello world"
	showType(x)
	showType(y)

}

func showType(t interface{}){
	fmt.Printf("The tipe of variable is %T and value is %v \n",t,t)
}