//Pointers

package main

func main(){
	// Memorie -> address -> Value
	a:=10
	println(a) //value
	println(&a) // address
	var pointer *int =&a  // using addres variable a 
	println(pointer)
	*pointer = 20 //change value of all variables that with same address of pointer
	println(a)
	b:=&a
	println(b) // print addres
	println(*b) //print value


}