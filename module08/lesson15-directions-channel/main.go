package main

func receive(name string, hello chan <- string) {
	hello <- name
}
func read(data <- chan string) {
	println("hello", <- data)
}

func main() {
	hello := make(chan string)
	go receive("John", hello)
	read(hello)
}
