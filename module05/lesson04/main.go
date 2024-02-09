//understanding go mod
package main
import "github.com/google/uuid"
func main()  {
	//using packeage uuid from google
	println(uuid.NewString())
}