package main

import (
	"fmt"
	_ "go_practice/p0016_init-func/bar"
	_ "go_practice/p0016_init-func/foo"
)

var global = convert()

func convert() int {
	return 100
}

func init() {
	global = 0
}

func main() {
	fmt.Println("global is", global)

	a := 121
	b := 1221
	c := 12321

	println(a, isPalindrome(a))
	println(b, isPalindrome(b))
	println(c, isPalindrome(c))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var a []int
	for i := x; i != 0; i /= 10 {
		a = append(a, i%10)
	}
	j := int8(len(a) - 1)
	for _, v := range a {
		if v != a[j] {
			return false
		}
		j--
	}
	return true
}
