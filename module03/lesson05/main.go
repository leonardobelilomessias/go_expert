// Context with value
package main

import (
	"context"
	"fmt"
)

func main()  {
	ctx:= context.Background()
	ctx = context.WithValue(ctx, "token", "password")
	bookHotel(ctx, "hibisco")
}
func bookHotel(cxt context.Context,name string){
	token := cxt.Value("token")
	fmt.Println(token)
}