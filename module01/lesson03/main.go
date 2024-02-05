//criation types
package main

const a  = "helo world"

var b bool  = true 
 var (
	c bool
	d int
	e float64
	f ID = 12
 )

type ID int

func main(){
	example := "example variable wite atribuition :=(this cannot be reatribuited with := only '=' is used)"
	println(example)
	println(b)
	b=false
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)

}