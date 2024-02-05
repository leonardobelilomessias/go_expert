// Package and modules

package main

import (
	"fmt"
 	"goexpert.com/module01/lesson21/mymath"
)

func main(){
	result := mymath.Sum(2,5)
	fmt.Println(" The result : ", result)
}