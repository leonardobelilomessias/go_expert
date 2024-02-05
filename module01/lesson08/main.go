// Functions
package main

import (
	"errors"
	"fmt"
)

func main(){
	 fmt.Println(sum(2,8))
	 fmt.Println(sum2(5,58))
	 fmt.Println(sum3(89,56))
	 //receiving value error
	 value, err:= sum3(123,200)
	 if err != nil{
		fmt.Println(err)
	 }
	 fmt.Println(value)
	}

func sum(a int, b int)int {
	return a + b
}

// Function with more then one value 
func sum2(a int, b int)(int, bool) {
	if(a+b >150){
		return a+b ,true
	}
	return a + b, false
}

// To handle error  in function 
func sum3(a int, b int)(int, error) {
	if(a+b >150){
		return a+b ,errors.New("sum is bigger then 150")
	}
	return a + b, nil
}