//  Custumization request

package main

import (
	"io"
	"net/http"
)

func main()  {
	// crreating http client 
	c := http.Client{}
	// creating request
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil{
		panic(err)
	}
	// add elements to request
	req.Header.Set("Accept","application/json")
	//enveloping request with client end doing requisition
	resp, err := c.Do(req)
	if err!= nil{
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}
	println(string(body))

}