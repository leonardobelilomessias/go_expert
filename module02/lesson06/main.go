// start with http
package main

import "net/http"


func main(){
	http.HandleFunc("/", FeatchCep)
	http.ListenAndServe(":8080", nil)
}

func  FeatchCep(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("hello world!"))
}