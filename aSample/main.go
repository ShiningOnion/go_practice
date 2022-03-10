package main

import "fmt"

func main() {
	println(HelloWorld("Leo"))
}

func HelloWorld(name string) string {
	return fmt.Sprintf("Hello, %s ", name)
}
