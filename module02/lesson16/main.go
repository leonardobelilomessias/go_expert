//templates composition
package main

import (
	"net/http"
	"text/template"
)


type Course struct{
	Name string 
	TimeWork int 
}
type Courses []Course
func main()  {
	templates :=[]string{
		"header.html",
		"content.html",
		"footer.html",
	}
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("content.html").ParseFiles(templates...))	
	err := t.Execute(w,  Courses{{"Golang",40},{"Nodejs",30}})
	if err!= nil{
		panic(err)
	}
	})
	http.ListenAndServe(":8080",nil)
	 

	}