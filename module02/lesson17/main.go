// templates composition
package main

import (
	"html/template"
	"os"
	"strings"
)


type Course struct{
	Name string 
	TimeWork int 
}
type Courses []Course

func Toupper(s string)string{
	return strings.ToUpper(s)
}

func main()  {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t:= template.New("content.html")
	t.Funcs(template.FuncMap{"Toupper":Toupper})
	t = template.Must(t.ParseFiles(templates...))	

	err := t.Execute(os.Stdout,  Courses{{"Golang",40},{"Nodejs",30}})
	if err!= nil{
		panic(err)
	}

	}