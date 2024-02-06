//make call to http
package main

import (
	"io"
	"net/http"
)

func main(){
	// get requisition
	req , err := http.Get("https://www.google.com")
	if err!= nil{
		panic(err)
	}
	//create a reader
	result, err:= io.ReadAll(req.Body)
	if err!= nil{

	}
	// print result
	println(string(result))
	//important close requisition
	req.Body.Close()

}