package main

import "os"

func main() {
	s, sep := "", ""
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		s += sep + arg
		sep = " "
		println("index:", i, "value:", arg)
	}
}
