//  Working with post requisition

package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main()  {
	c := http.Client{}
	// is needed to get transform json to byte buffer 
	jsonvar := bytes.NewBuffer([]byte(`{name:"john}`))
	resp, err := c.Post("http://google.com","application/json", jsonvar)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}