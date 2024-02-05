// Slices
package main

import "fmt"

func main(){
	x := []int{2,55,8,55,5}
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
	fmt.Printf("len=%d cap=%d %v\n", len(x[:0]), cap(x[:0]), x[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(x[:2]), cap(x[:2]), x[:2])
	fmt.Printf("len=%d cap=%d %v\n", len(x[2:]), cap(x[2:]), x[2:])
	x = append(x, 57)
	fmt.Printf("len=%d cap=%d %v\n", len(x[2:]), cap(x[2:]), x[2:])


}