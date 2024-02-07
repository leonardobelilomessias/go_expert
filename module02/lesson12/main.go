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
	 tmp := template.New("CouseTemplate")
	 tmp,_ = tmp.Parse("Curse:{{.Name}} . Time to Work {{.TimeWork}}")
	 err:= tmp.Execute(os.Stdout, course)
	 if err!= nil{
		panic(err)
	 }

	}