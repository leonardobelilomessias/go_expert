//get start with templates

package main

import (
	"os"
	"text/template"
)


type Course struct{
	Name string 
	TimeWork int 
}
type Courses []Course
func main()  {
	
	 t := template.Must(template.New("template.html").ParseFiles("template.html"))	
	 err := t.Execute(os.Stdout,  Courses{{"Golang",40},{"Nodejs",30}})
	 if err!= nil{
		panic(err)
	 }

	}