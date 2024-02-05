package main

import(
	"github.com/google/uuid"
)
func main(){
	myId  :=  uuid.New()
	println(myId.ID())
}