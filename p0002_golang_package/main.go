package main

import (
	"fmt"

	"github.com/ShiningOnion/go_practice/p0002_golang_package/controller"
)

func main() {
	println("practice package from github")

	hi := controller.HelloWorld("Leo")
	fmt.Println(hi)
}
