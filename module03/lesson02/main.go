// Basic concepts context
package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx := context.Background()
	// the context timeout will execute done when finished time
	ctx, cancel := context.WithTimeout(ctx, time.Second * 6)
	//if  the timeout be less then time execution wil send done
	//ctx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()
	bookHotel(ctx) 
}

func bookHotel(ctx context.Context){
	select{
	case <- ctx.Done():
		fmt.Println("Hotal booking cancelled. Timeout exceded.")
		return
	case <- time.After(2* time.Second):
			fmt.Println("Hotel Booked.")
			return
	}
}