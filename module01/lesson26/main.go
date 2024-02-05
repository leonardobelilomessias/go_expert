//conditions
package main

func main(){
	a:=1
	b:=3
	c :=4
	if a >b && c > a{
		println(a)
	}else{
		println(b)
	}

	switch c {
	case 1:
		println(a)
	case 3:
		println(b)
	default:
		println("no value")
	}
}