// server mux
package main

import (
	"net/http"
)


func main(){
 mux:= http.NewServeMux()
 mux.HandleFunc("/", HomeHandler)
 mux.Handle("/blog",blog{title: "My Blog"})
 http.ListenAndServe(":8080", mux)
}

func HomeHandler (response http.ResponseWriter, request *http.Request)  {
	response.Write([]byte("Hellow world"))
 }

 type blog struct{
	title string
 }

 func (b blog) ServeHTTP(response http.ResponseWriter, request *http.Request)  {
	response.Write([]byte(b.title))
 }