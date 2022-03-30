package main

import (
	"fmt"
	"go_practice/self/p0011_cross-build/hello"
	"go_practice/self/p0011_cross-build/hello2"
	"go_practice/self/p0011_cross-build/hello3"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(hello2.Hello())
	fmt.Println(hello3.Hello())
}
