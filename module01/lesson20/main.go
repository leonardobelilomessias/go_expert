//Generics
package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func sum[T Number] (m map[string] T) T {
	var sum  T 
	for _, v := range m {
		sum += v
	}
	return sum
}

func compare[T comparable] (a T, b T) bool{
	if a == b { 
		return true
	}
	return false
}

func main(){
	m:=map[string] int{ "john":3000, "peter":2000}
	m2:=map[string] float64{ "john":3000.2, "peter":2000.4}
	m3:=map[string] MyNumber{ "john":3000, "peter":2000}

	println(sum(m))
	println(sum(m2))
	println(sum(m3))
	println(compare(10,10))

}