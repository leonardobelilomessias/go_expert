// contex side client
package main

import (
	"log"
	"net/http"
	"time"
)

func main()  {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log.Println("Started request")
	defer log.Println("Finished request")
	select{
	case <- time.After(5* time.Second):
		w.Write([]byte("Request processed successfuly"))
	case<- ctx.Done():
		log.Println("Request canceled by cliente")
		http.Error(w,"Request canceled by cliente", http.StatusRequestTimeout)
	}
}