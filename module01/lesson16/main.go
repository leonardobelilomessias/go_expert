// When should use pointers

package main

func sum(a int, b int) int{
	a = 50
	return a + b 
}

func sum2(a *int, b *int) int{
	*a = 50
	return *a + *b 
}
func main(){
	myvar1 :=10
	myvar2:=20
	sum(myvar1,myvar2)
	println(sum(myvar1,myvar2)) 
	println(myvar1)// variables no change cause we just pass the value dont reference memorie
	println(sum2(&myvar1,&myvar2))
	println(myvar1) //chan value variable by the pointer
}
