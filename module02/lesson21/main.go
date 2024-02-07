// Working with http using contect

package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main()  {
	 ctx := context.Background()
	 //exced time on microsecond 
	 ctx, cancel:= context.WithTimeout(ctx, time.Microsecond)
	 defer cancel()
	 req, err:= http.NewRequestWithContext(ctx,"GET", "http://google.com", nil)
	 if err!= nil{
		panic(err)
	 }
	 resp, err := http.DefaultClient.Do(req)
	 if err != nil{
		panic(err)
	 }
	 defer resp.Body.Close()
	 body, err:= io.ReadAll(resp.Body)
	 if err!= nil {
		 panic(err)
	 }
	 println(string(body))
}
