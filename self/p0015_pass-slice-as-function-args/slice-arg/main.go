package main

import "fmt"

// pass slice as function argument
func modify(foo []string) []string {
	foo[1] = "1"
	fmt.Println("modify foo", foo)
	return foo
}

func main() {
	foo := []string{"a", "b"}
	fmt.Println("before foo:", foo)
	modify(foo)
	fmt.Println("after foo:", foo)
}
