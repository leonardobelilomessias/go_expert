//Running Arrays

package main

import "fmt"

func main(){
	 var  myarray[3]int

	 myarray[0] = 55
	 myarray[1] = 75
	 myarray[2] = 22
	 fmt.Println(myarray)
	 fmt.Println(len(myarray))
	 
	 for i := 0; i < len(myarray); i++ {
		fmt.Print(myarray[i], " ")
	 }

	 println()

	  for i,v := range myarray{
		fmt.Printf("value of indice %d é  valor é %d \n",i,v)
	  }
}