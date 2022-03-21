package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hi, Leo Docker Build")

	// try changing the value of this url
	res, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Status)
}
