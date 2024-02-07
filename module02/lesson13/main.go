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

func main()  {
	 course := Course{"Golang",40}
	 t := template.Must(template.New(" CourseTamplate").Parse("Curse:{{.Name}} . Time to Work {{.TimeWork}}"))	
	 err := t.Execute(os.Stdout, course)
	 if err!= nil{
		panic(err)
	 }

	}