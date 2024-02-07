//  Http client timeout

package main

import (
	"io"
	"net/http"
	"time"
)

func main()  {
	c := http.Client{
		//timeout throw by microsecon, if change to second will works 
		Timeout: time.Microsecond,
	}
	resp, err := c.Get("http://google.com")
	if err != nil{
		panic(err)
	}
	body, err:= io.ReadAll(resp.Body)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	println(string(body))
}