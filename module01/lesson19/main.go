// Type assertion

package main

func main(){
	var myvar  interface{} = "Jonh Due"
	println(myvar.(string)) // to afirme that it is string
	res, ok := myvar.(int) // wrong result 
	println(res)
	println(ok)
	res2 :=myvar.(int)
	print(res2)
}