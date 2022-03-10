package main

import "fmt"

func main() {
	a := 100
	b := 50
	println(HelloWorld("Leo"))

	if a >= b {
		fmt.Println(a, " >= ", b)
	}
}

func HelloWorld(name string) string {
	return fmt.Sprintf("Hello, %s ", name)
}
