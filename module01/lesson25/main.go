//Loops

package main

func main(){
 for i := 0; i < 10; i++ {
	print(i)
 }
 
 println()
 numbers :=[]string{"one", "two", "twe"}
 for  key, value := range numbers{
	print(key,value)
 }
 println()

 i:=0
 for i<10{ //looks like while
	println(i)
	i++
 }
}