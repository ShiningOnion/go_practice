package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]*[]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, values := range counts {
		n := len(*values)
		if n > 1 {
			fmt.Printf("%d\t%s\t%+v\n", n, line, *values)
		}
	}
}

func countLines(f *os.File, counts map[string]*[]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if val, ok := counts[input.Text()]; ok {
			//do something here
			*val = append(*val, f.Name())
			// println(len(*val), input.Text(), f.Name())
		} else {
			counts[input.Text()] = &[]string{f.Name()}
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
