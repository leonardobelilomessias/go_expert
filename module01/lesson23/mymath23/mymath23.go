package mymath23

// only elements with Capitalize are exported from package

func Sum[T int | float64](a,b T) T{
	return a + b
}
var A int = 10
var b int = 12