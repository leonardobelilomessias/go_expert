// handle file
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//create a file
	file, err := os.Create("arquivo.txt")
	if err != nil{
		panic(err)
	}
	//write a string in the  file
	size, err := file.WriteString("Hello, world! ")
	fmt.Printf(" File created sizr %d\n",size )
	file.Close()
	//reade file
	readeFile , err:= os.ReadFile("arquivo.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(readeFile))
	// open file
	file2, err := os.Open("arquivo.txt")
	if err!= nil{
		panic(err)
	}
	//create a buffer
	reader := bufio.NewReader(file2)
	// chose the size buffer
	buffer := make([]byte,3)

	// iterating to read buffer
	for{
		n, err:= reader.Read(buffer)
		if err !=nil{
			break
		}
	// join chuncks
		fmt.Println(string(buffer[:n]))
	}
	// delete file
	err = os.Remove("arquivo.txt")
	if err!= nil{
		panic(err)
	}
}