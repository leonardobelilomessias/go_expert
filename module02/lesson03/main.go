//defer
package main



func main(){
	// defer create a delay at line that will execute at end code execution.
	println("First line")
	defer println("Second line")
	println("Third line")
}