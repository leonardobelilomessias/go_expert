// Functions variables
package main

import (

	"fmt"
)

func main(){
	fmt.Println(sum(2,3,66,8,77))
	}

func sum(numbers  ...int)int {
	total := 0

	for _, number := range numbers{
		total += number
	}
	return total
}

